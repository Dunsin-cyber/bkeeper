package handlers

import (
	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/cmd/api/services"
	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdatePasswordHandler(c echo.Context) error {
	userService := services.NewUserSrvice(h.DB)
	//1 get user from middleware
	user, ok := c.Get("user").(models.UserModel)
	if !ok {
		return common.SendInternalServerErrorResponse(c, "User authentication failed")
	}
	//2bind the request body
	payload := new(requests.ChangePasswordRequest)
	if err := h.BindRequest(c, payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	//3 validate the data sent by client
	validationErrs := h.ValidateBodyRequest(c, *payload)
	if validationErrs != nil {
		return common.SendFailedValidationResponse(c, validationErrs)
	}

	//4 check if current password matches the already existing pw in the database
	validCurrentPassword := common.ComparePasswordHash(payload.CurrentPassword, user.Password)

	if !validCurrentPassword {
		errMsg := "the supplied password does not match your current password"
		return common.SendUnauthorizedResponse(c, &errMsg)
	}

	err := userService.ChangeAuthUserPassword(&user, payload.Password)
	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "Password change successful", nil)

}
