package handlers

import (
	"net/http"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	//bind the request body
	payload := new(requests.RegisterUserRequest)
	if err := c.Bind(payload); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	//validation
	validationErrs := h.ValidateBodyRequest(c, *payload)

	 if validationErrs != nil {
		return c.JSON(http.StatusBadRequest, validationErrs)
	 }

	//vaidate what we binded
	return c.JSON(http.StatusOK, "validation successful")

}
