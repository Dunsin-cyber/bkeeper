package handlers

import (
	"errors"
	"fmt"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/cmd/api/services"
	"github.com/Dunsin-cyber/bkeeper/common"
	"github.com/Dunsin-cyber/bkeeper/internal/app_errors"
	"github.com/Dunsin-cyber/bkeeper/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) ListCategories(c echo.Context) error {
	categoryService := services.NewCategoryService(h.DB)
	categories, err := categoryService.List(h.DB)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "Ok", categories)

}

func (h *Handler) CreateCategory(c echo.Context) error {
	//bind request
	payload := new(requests.CreateCatagoryRequest)
	if err := h.BindRequest(c, payload); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	//validation
	validationErrs := h.ValidateBodyRequest(c, *payload)

	if validationErrs != nil {
		return common.SendFailedValidationResponse(c, validationErrs)
	}

	categoryService := services.NewCategoryService(h.DB)
	categories, err := categoryService.Create(*payload)

	if err != nil {
		return common.SendInternalServerErrorResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "Category created", categories)

}

func (h *Handler) DeleteCategory(c echo.Context) error {
	_, ok := c.Get("user").(models.UserModel)
	if !ok {
		errMsg := "User authentication failed"
		return common.SendUnauthorizedResponse(c, &errMsg)
	}
	var categoryId requests.IDParamRequest
	if err := h.BindRequest(c, &categoryId); err != nil {
		return common.SendBadRequestResponse(c, err.Error())
	}

	fmt.Printf("The ID is: %d\n", categoryId.ID)

	categoryService := services.NewCategoryService(h.DB)
	err := categoryService.DeleteById(categoryId.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewNotFoundError(err.Error())
		}
		return common.SendNotFoundResponse(c, err.Error())
	}

	return common.SendSuccessResponse(c, "Category deleted", nil)

}
