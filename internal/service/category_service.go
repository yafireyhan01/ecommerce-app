package service

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(guid string) error
	GetCategoryByGuid(guid string) (*models.Category, error)
	GetAllCategories(page, limit int) ([]models.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.categoryRepo.Create(category)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *categoryService) DeleteCategory(guid string) error {
	return s.categoryRepo.Delete(guid)
}

func (s *categoryService) GetCategoryByGuid(guid string) (*models.Category, error) {
	return s.categoryRepo.FindByID(guid)
}

func (s *categoryService) GetAllCategories(page, limit int) ([]models.Category, error) {
	offset := (page - 1) * limit
	return s.categoryRepo.FindAll(offset, limit)
}
