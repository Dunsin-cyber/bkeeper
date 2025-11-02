package handlers

import "gorm.io/gorm"

type Handler struct {
	// Add fields for dependencies like database, logger, etc.
	DB *gorm.DB
}
