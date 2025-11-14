package handlers

import (
	"encoding/base64"
	"errors"
	"net/url"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/cmd/api/services"
	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/mailer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) ForgotPasswordHandler(c echo.Context) error {
	//bind the request body
	payload := new(requests.ForgotPasswordRequest)
	if err := h.BindRequest(c, payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	//validation
	validationErrs := h.ValidateBodyRequest(c, *payload)

	if validationErrs != nil {
		return common.SendFailedValidationResponse(c, validationErrs)
	}

	userService := services.NewUserSrvice(h.DB)
	appTokenService := services.NewAppTokenService(h.DB)
	// check if user(email) exists
	retrievedUser, err := userService.GetUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.SendNotFoundResponse(c, "Record not found, register with this email")
		}
		return common.SendInternalServerErrorResponse(c, "An error occured, try again later")

	}

	token, err := appTokenService.GenerateResetPasswordToken(*retrievedUser)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, "An error occured, please try again later")
	}

	encodedEmail := base64.RawURLEncoding.EncodeToString([]byte(retrievedUser.Email))

	frontendUrl, err := url.Parse(payload.FrontendURL)
	if err != nil {
		return common.SendBadRequestResponse(c, "invalid frontend URL")
	}

	query := url.Values{}
	query.Set("email", encodedEmail)
	query.Set("token", token.Token)
	frontendUrl.RawQuery = query.Encode()

	// send mail to user that contains the reset password token
	mailData := mailer.EmailData{
		Subject: "Request Password Reset",
		Meta: struct {
			Token       string
			FrontendURL string
		}{
			Token:       token.Token,
			FrontendURL: frontendUrl.String(),
		},
	}
	err = h.Mailer.Send(payload.Email, "forgot-password.html", mailData)
	if err != nil {
		h.Logger.Error("Failed to send forgot password email: ", err)
	}

	return common.SendSuccessResponse(c, "Forgot password email sent", nil)
}

func (h *Handler) ResetPasswordHandler(c echo.Context) error {
	//bind the request body
	payload := new(requests.ResetPasswordRequest)
	if err := h.BindRequest(c, payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	//validation
	validationErrs := h.ValidateBodyRequest(c, *payload)

	if validationErrs != nil {
		return common.SendFailedValidationResponse(c, validationErrs)
	}

	email, err := base64.RawURLEncoding.DecodeString(payload.Meta)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, "An error occured, please try again later")
	}

	userService := services.NewUserSrvice(h.DB)
	appTokenService := services.NewAppTokenService(h.DB)

	retrievedUser, err := userService.GetUserByEmail(string(email))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.SendNotFoundResponse(c, "Invalid password reset token")
		}
	}

	appToken, err := appTokenService.ValidateResetPasswordToken(*retrievedUser, payload.Token)

	if err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	err = userService.ChangeAuthUserPassword(retrievedUser, payload.Password)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	appTokenService.InvalidateToken(retrievedUser.ID, *appToken)

	return common.SendSuccessResponse(c, "Reset password completed", nil)

}
