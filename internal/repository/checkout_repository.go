package repository

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"gorm.io/gorm"
)

type CheckoutRepository interface {
	Create(checkout *models.Checkout) error
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
