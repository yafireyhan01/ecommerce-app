package repository

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"gorm.io/gorm"
)

type CheckoutRepository interface {
	Create(checkout *models.Checkout) error
	FindByID(guid string) (*models.Checkout, error)
	Update(category *models.Checkout) error
}

type checkoutRepository struct {
	db *gorm.DB
}

func NewCheckoutRepository(db *gorm.DB) CheckoutRepository {
	return &checkoutRepository{db}
}

func (r *checkoutRepository) Create(checkout *models.Checkout) error {
	return r.db.Create(checkout).Error
}

func (r *checkoutRepository) FindByID(guid string) (*models.Checkout, error) {
	var checkout models.Checkout
	err := r.db.Where("guid = ? AND deleted_at IS NULL", guid).First(&checkout).Error
	return &checkout, err
}

func (r *checkoutRepository) Update(category *models.Checkout) error {
	return r.db.Save(category).Error
}
