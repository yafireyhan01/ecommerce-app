package repository

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"gorm.io/gorm"
	"time"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(guid string) error
	FindByID(guid string) (*models.Category, error)
	FindAll(offset, limit int) ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(guid string) error {
	return r.db.Model(&models.Category{}).Where("guid = ?", guid).Update("deleted_at", time.Now()).Error
}

func (r *categoryRepository) FindByID(guid string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("guid = ? AND deleted_at IS NULL", guid).First(&category).Error
	return &category, err
}

func (r *categoryRepository) FindAll(offset, limit int) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Model(&models.Category{}).Select("guid, name").Where("deleted_at IS NULL").Offset(offset).Limit(limit).Find(&categories).Error
	return categories, err
}
