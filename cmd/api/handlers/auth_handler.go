package handlers

import (
	"net/http"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	//bind the request body
	payload := new(requests.RegisterUserRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	c.Logger().Print(payload)

	//vaidate what we binded
	return c.JSON(http.StatusOK, payload)

}
