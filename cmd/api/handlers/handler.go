package handlers

import (
	"errors"

	"github.com/Dunsin-cyber/bkeeper/internal/mailer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	// Add fields for dependencies like database, logger, etc.
	DB     *gorm.DB
	Logger echo.Logger
	Mailer mailer.Mailer
}

func (h *Handler) BindRequest(c echo.Context, payload interface{}) error {

	if err := c.Bind(payload); err != nil {
		c.Logger().Error(err)
		return errors.New("failed to bind request, make sure you are sending a valid payload")
	}

	return nil
}
