package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/service"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
	"strconv"
)

type CategoryHandler struct {
	categoryService service.CategoryService
	validate        *validator.Validate
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		validate:        validator.New(),
	}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var input struct {
		Name string `json:"name" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	category := &models.Category{
		Guid: utils.NewUUID(),
		Name: input.Name,
	}

	if err := h.categoryService.CreateCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create category"})
	}

	return c.JSON(fiber.Map{
		"guid": category.Guid,
		"name": category.Name,
	})
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	guid := c.Params("guid")
	var input struct {
		Name string `json:"name" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	category, err := h.categoryService.GetCategoryByGuid(guid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	category.Name = input.Name

	if err := h.categoryService.UpdateCategory(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update category"})
	}

	return c.JSON(fiber.Map{
		"guid": category.Guid,
		"name": category.Name,
	})
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	guid := c.Params("guid")

	if err := h.categoryService.DeleteCategory(guid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete category"})
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}

func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	guid := c.Params("guid")

	category, err := h.categoryService.GetCategoryByGuid(guid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	return c.JSON(fiber.Map{
		"guid": category.Guid,
		"name": category.Name,
	})
}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
	}

	limit := 10

	categories, err := h.categoryService.GetAllCategories(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve categories"})
	}

	response := make([]fiber.Map, len(categories))
	for i, category := range categories {
		response[i] = fiber.Map{
			"guid": category.Guid,
			"name": category.Name,
		}
	}

	return c.JSON(response)
}
