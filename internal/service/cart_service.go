package service

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
)

type CartService interface {
	CreateCart(cart *models.Cart) error
	UpdateCart(cart *models.Cart) error
	DeleteCart(guid string) error
	GetCartByGuid(guid string) (*models.Cart, error)
	GetAllCarts(page, limit int) ([]models.Cart, error)
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{cartRepo}
}

func (s *cartService) CreateCart(cart *models.Cart) error {
	return s.cartRepo.Create(cart)
}

func (s *cartService) UpdateCart(cart *models.Cart) error {
	return s.cartRepo.Update(cart)
}

func (s *cartService) DeleteCart(guid string) error {
	return s.cartRepo.Delete(guid)
}

func (s *cartService) GetCartByGuid(guid string) (*models.Cart, error) {
	return s.cartRepo.FindByID(guid)
}

func (s *cartService) GetAllCarts(page, limit int) ([]models.Cart, error) { // Implement this method
	offset := (page - 1) * limit
	return s.cartRepo.FindAll(offset, limit)
}
