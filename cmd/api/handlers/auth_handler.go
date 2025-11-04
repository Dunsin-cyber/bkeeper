package handlers

import (

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/labstack/echo/v4"
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

	return common.SendSuccessResponse(c,"User resgistration successful", nil)

}
