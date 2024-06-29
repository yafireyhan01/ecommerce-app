package repository

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"gorm.io/gorm"
	"time"
)

type ProductRepository interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(guid string) error
	FindByID(guid string) (*models.Product, error)
	FindAll(offset, limit int) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(guid string) error {
	return r.db.Model(&models.Product{}).Where("guid = ?", guid).Update("deleted_at", time.Now()).Error
}

func (r *productRepository) FindByID(guid string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("guid = ? AND deleted_at IS NULL", guid).First(&product).Error
	return &product, err
}

func (r *productRepository) FindAll(offset, limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Model(&models.Product{}).Select("guid, category_guid, name, description, price, stock_qty").Where("deleted_at IS NULL").Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}
