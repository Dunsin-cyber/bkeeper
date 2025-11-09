package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	// Add fields for dependencies like database, logger, etc.
	DB *gorm.DB
	Logger  echo.Logger
}
