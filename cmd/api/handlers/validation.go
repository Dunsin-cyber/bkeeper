package handlers

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationError struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

func (h *Handler) ValidateBodyRequest(c echo.Context, payload interface{}) []*ValidationError {

	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	var errors []*ValidationError
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


			switch condition {
			case "required":
				errMsg = keyToTitleCase + " field is required"
			case "email":
				errMsg = keyToTitleCase + " must be a valid email address"
			case "min":
				errMsg = keyToTitleCase + " must be at least " + validationErr.Param() + " characters long"
			case "max":
				errMsg = keyToTitleCase + " must be at most " + validationErr.Param() + " characters long"
			}


			currentValidationError := &ValidationError{
				Error:      errMsg,
				Key:       key,
				Condition: condition,
			}
			errors = append(errors, currentValidationError)
		}
	}

	return errors
}
