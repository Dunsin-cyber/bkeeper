package handlers

import (
	"errors"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/cmd/api/services"
	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/mailer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	//bind the request body
	payload := new(requests.RegisterUserRequest)
	if err := c.Bind(payload); err != nil {
		c.Logger().Error(err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	//validation
	validationErrs := h.ValidateBodyRequest(c, *payload)

	if validationErrs != nil {
		return common.SendFailedValidationResponse(c, validationErrs)
	}

	userService := services.NewUserSrvice(h.DB)
	// check if user(email) exists
	_, err := userService.GetUserByEmail(payload.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "Email has already been taken")
	}
	//create a user
	result, err := userService.RegisterUser(*payload)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	//send a welcome message to the user
	mailData := mailer.EmailData{
		Subject: "Welcome to Bkeeper Finance",
		Meta: struct {
			FirstName string
			LoginLink string
		}{
			FirstName: *result.FirstName,
			LoginLink: "#",
		},
	}
	err = h.Mailer.Send(payload.Email, "welcome.html", mailData)
	if err != nil {
		h.Logger.Error("Failed to send welcome email: ", err)
	}

	//send response
	return common.SendSuccessResponse(c, "User resgistration successful", result)

}

func (h *Handler) LoginHandler(c echo.Context) error {
	userService := services.NewUserSrvice(h.DB)
	//1 bind our data
	payload := new(requests.LoginUserRequest)
	if err := c.Bind(payload); err != nil {
		c.Logger().Error(err)
		return common.SendBadRequestResponse(c, err.Error())
	}

	//2 validate the data sent by client
	validationErrs := h.ValidateBodyRequest(c, *payload)
	if validationErrs != nil {
		return common.SendFailedValidationResponse(c, validationErrs)
	}
	//3 if user with supplied email exists
	userRetrieved, err := userService.GetUserByEmail(payload.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.SendBadRequestResponse(c, "invalid email or password")
	}
	//4 compare the client pw with the hashed pw
	validPassword := common.ComparePasswordHash(payload.Password, userRetrieved.Password)
	if !validPassword {
		return common.SendBadRequestResponse(c, "invalid email or password")

	}
	accessToken, refreshToken, err := common.GenerateJWT(*userRetrieved)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}
	//5 send a response
	return common.SendSuccessResponse(c, "User logged in", map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          userRetrieved,
	})

}
