package service

import (
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(guid string) error
	GetProductByGuid(guid string) (*models.Product, error)
	GetAllProducts(page, limit int) ([]models.Product, error)
	GetProductsByCategoryGuid(categoryGuid string, page, limit int) ([]models.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.productRepo.Create(product)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(guid string) error {
	return s.productRepo.Delete(guid)
}

func (s *productService) GetProductByGuid(guid string) (*models.Product, error) {
	return s.productRepo.FindByID(guid)
}

func (s *productService) GetAllProducts(page, limit int) ([]models.Product, error) {
	offset := (page - 1) * limit
	return s.productRepo.FindAll(offset, limit)
}

func (s *productService) GetProductsByCategoryGuid(categoryGuid string, page, limit int) ([]models.Product, error) {
	offset := (page - 1) * limit
	return s.productRepo.FindByCategoryGuid(categoryGuid, offset, limit)
}
