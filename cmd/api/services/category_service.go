package services

import (
	"errors"
	"strings"

	"github.com/Dunsin-cyber/bkeeper/cmd/api/requests"
	"github.com/Dunsin-cyber/bkeeper/internal/app_errors"
	"github.com/Dunsin-cyber/bkeeper/internal/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (c CategoryService) List(db *gorm.DB) ([]*models.CategoryModel, error) {
	var categories []*models.CategoryModel

	result := c.db.Find(&categories)
	if result.Error != nil {
		return nil, errors.New("failed to fetch categories")
	}

	return categories, nil

}

func (c CategoryService) Create(data requests.CreateCatagoryRequest) (*models.CategoryModel, error) {

	slug := strings.ToLower(data.Name)
	slug = strings.Replace(slug, " ", "_", -1)

	categoryCreated := &models.CategoryModel{
		Slug:     slug,
		Name:     data.Name,
		IsCustom: data.IsCustom,
	}

	result := c.db.Where(models.CategoryModel{Slug: slug}).FirstOrCreate(categoryCreated)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return categoryCreated, nil
	}

	if result.Error != nil {
		return nil, errors.New("failed to create a new category")
	}

	return categoryCreated, nil

}

func (c CategoryService) GetById(id uint) (*models.CategoryModel, error) {
	var category *models.CategoryModel

	result := c.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_errors.NewNotFoundError("Category not found")
		}
		return nil, errors.New("failed to fetch category")

	}
	return category, nil
}

func (c CategoryService) DeleteById(id uint) error {
	var category *models.CategoryModel

	category, err := c.GetById(id)
	if err != nil {
		return err
	}
	c.db.Delete(category)
	return nil
}
