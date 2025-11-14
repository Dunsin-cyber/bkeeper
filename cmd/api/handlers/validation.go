package handlers

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ValidateBodyRequest(c echo.Context, payload interface{}) []*common.ValidationError {

	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	var errors []*common.ValidationError
	err := validate.Struct(payload)

	validationErrors, ok := err.(validator.ValidationErrors)

	if ok {
		reflected := reflect.ValueOf(payload)

		for _, validationErr := range validationErrors {
			field, _ := reflected.Type().FieldByName(validationErr.StructField())

			key := field.Tag.Get("json")
			if key == " " {
				key = strings.ToLower(validationErr.StructField())
			}

			condition := validationErr.Tag()

			keyToTitleCase := strings.Replace(key, "_", " ", -1)
			errMsg := keyToTitleCase + " field is " + condition
			param := validationErr.Param()

			switch condition {
			case "required":
				errMsg = keyToTitleCase + " field is required"
			case "email":
				errMsg = keyToTitleCase + " must be a valid email address"
			case "min":
				if _, err := strconv.Atoi(param); err == nil {

					errMsg = keyToTitleCase + " must be at least " + param + " characters long"
				} else {
					errMsg = keyToTitleCase + " must be at least " + param

				}
			case "max":
				if _, err := strconv.Atoi(param); err == nil {

					errMsg = keyToTitleCase + " must be at most " + param + " characters long"
				} else {
					errMsg = keyToTitleCase + " must be at most " + param

				}
			case "eqfield":
				errMsg = keyToTitleCase + " must be equal to " + strings.ToLower(param)
			case "url":
				errMsg = keyToTitleCase + " must be a valid url"
			}

			currentValidationError := &common.ValidationError{
				Error: errMsg,
				Key:   key,
			}
			errors = append(errors, currentValidationError)
		}
	}

	return errors
}
