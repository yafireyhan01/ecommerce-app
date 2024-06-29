package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/service"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
	"strconv"
)

type ProductHandler struct {
	productService service.ProductService
	validate       *validator.Validate
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		validate:       validator.New(),
	}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var input struct {
		CategoryGuid string  `json:"category_guid" validate:"required"`
		Name         string  `json:"name" validate:"required"`
		Description  string  `json:"description"`
		Price        float64 `json:"price" validate:"required"`
		StockQty     int     `json:"stock_qty" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	categoryGuid, err := uuid.FromString(input.CategoryGuid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category GUID"})
	}

	product := &models.Product{
		Guid:         utils.NewUUID(),
		CategoryGuid: categoryGuid,
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		StockQty:     input.StockQty,
	}

	if err := h.productService.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create product"})
	}

	return c.JSON(fiber.Map{
		"guid":          product.Guid,
		"category_guid": product.CategoryGuid,
		"name":          product.Name,
		"description":   product.Description,
		"price":         product.Price,
		"stock_qty":     product.StockQty,
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	guid := c.Params("guid")
	var input struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"required"`
		StockQty    int     `json:"stock_qty" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := h.productService.GetProductByGuid(guid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.StockQty = input.StockQty

	if err := h.productService.UpdateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update product"})
	}

	return c.JSON(fiber.Map{
		"guid":        product.Guid,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock_qty":   product.StockQty,
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	guid := c.Params("guid")

	if err := h.productService.DeleteProduct(guid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete product"})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	guid := c.Params("guid")

	product, err := h.productService.GetProductByGuid(guid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(fiber.Map{
		"guid":          product.Guid,
		"category_guid": product.CategoryGuid,
		"name":          product.Name,
		"description":   product.Description,
		"price":         product.Price,
		"stock_qty":     product.StockQty,
	})
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
	}

	limit := 10

	products, err := h.productService.GetAllProducts(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve products"})
	}

	response := make([]fiber.Map, len(products))
	for i, product := range products {
		response[i] = fiber.Map{
			"guid":          product.Guid,
			"category_guid": product.CategoryGuid,
			"name":          product.Name,
			"description":   product.Description,
			"price":         product.Price,
			"stock_qty":     product.StockQty,
		}
	}

	return c.JSON(response)
}
