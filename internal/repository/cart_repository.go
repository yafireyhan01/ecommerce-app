package repository

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"gorm.io/gorm"
	"time"
)

type CartRepository interface {
	Create(cart *models.Cart) error
	Update(cart *models.Cart) error
	Delete(guid string) error
	FindByID(guid string) (*models.Cart, error)
	FindAll(offset, limit int) ([]models.Cart, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) Create(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) Update(cart *models.Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepository) Delete(guid string) error {
	return r.db.Model(&models.Cart{}).Where("guid = ?", guid).Update("deleted_at", time.Now()).Error
}

func (r *cartRepository) FindByID(guid string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("guid = ?", guid).First(&cart).Error
	return &cart, err
}

func (r *cartRepository) FindAll(offset, limit int) ([]models.Cart, error) {
	var categories []models.Cart
	err := r.db.Model(&models.Cart{}).Select("guid, product_guid, qty").Where("deleted_at IS NULL").Offset(offset).Limit(limit).Find(&categories).Error
	return categories, err
}
