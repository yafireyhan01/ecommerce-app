package service

import (
	"github.com/gofrs/uuid"
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
)

type CheckoutService interface {
	CreateCheckout(checkout *models.Checkout) error
	CalculateTotalPrice(cartGuid uuid.UUID) (float64, error)
}

type checkoutService struct {
	checkoutRepo repository.CheckoutRepository
	cartRepo     repository.CartRepository
	productRepo  repository.ProductRepository
}

func NewCheckoutService(checkoutRepo repository.CheckoutRepository, cartRepo repository.CartRepository, productRepo repository.ProductRepository) CheckoutService {
	return &checkoutService{
		checkoutRepo: checkoutRepo,
		cartRepo:     cartRepo,
		productRepo:  productRepo,
	}
}

func (s *checkoutService) CreateCheckout(checkout *models.Checkout) error {
	return s.checkoutRepo.Create(checkout)
}

func (s *checkoutService) CalculateTotalPrice(cartGuid uuid.UUID) (float64, error) {
	cart, err := s.cartRepo.FindByID(cartGuid.String())
	if err != nil {
		return 0, err
	}

	product, err := s.productRepo.FindByID(cart.ProductGuid.String())
	if err != nil {
		return 0, err
	}

	totalPrice := float64(cart.Qty) * product.Price
	return totalPrice, nil
}
