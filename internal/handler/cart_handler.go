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

type CartHandler struct {
	cartService service.CartService
	validate    *validator.Validate
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
		validate:    validator.New(),
	}
}

func (h *CartHandler) CreateCart(c *fiber.Ctx) error {
	userGuid := c.Locals("UserGuid").(string)
	var input struct {
		ProductGuid string `json:"product_guid" validate:"required"`
		Qty         int    `json:"qty" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	cart := &models.Cart{
		Guid:        utils.NewUUID(),
		UserGuid:    uuid.FromStringOrNil(userGuid),
		ProductGuid: uuid.FromStringOrNil(input.ProductGuid),
		Qty:         input.Qty,
	}

	if err := h.cartService.CreateCart(cart); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create cart"})
	}

	return c.JSON(fiber.Map{
		"guid":         cart.Guid,
		"user_guid":    cart.UserGuid,
		"product_guid": cart.ProductGuid,
		"qty":          cart.Qty,
	})
}

func (h *CartHandler) GetAllCarts(c *fiber.Ctx) error { // Add this method
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit"})
	}

	carts, err := h.cartService.GetAllCarts(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve carts"})
	}

	response := make([]fiber.Map, len(carts))
	for i, cart := range carts {
		response[i] = fiber.Map{
			"guid":         cart.Guid,
			"user_guid":    cart.UserGuid,
			"product_guid": cart.ProductGuid,
			"qty":          cart.Qty,
		}
	}

	return c.JSON(response)
}

func (h *CartHandler) UpdateCart(c *fiber.Ctx) error {
	guid := c.Params("guid")
	var input struct {
		Qty int `json:"qty" validate:"required,min=1"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	cart, err := h.cartService.GetCartByGuid(guid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cart not found"})
	}

	cart.Qty = input.Qty

	if err := h.cartService.UpdateCart(cart); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update cart"})
	}

	return c.JSON(fiber.Map{
		"guid":         cart.Guid,
		"user_guid":    cart.UserGuid,
		"product_guid": cart.ProductGuid,
		"qty":          cart.Qty,
	})
}

func (h *CartHandler) DeleteCart(c *fiber.Ctx) error {
	guid := c.Params("guid")

	if err := h.cartService.DeleteCart(guid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete cart"})
	}

	return c.JSON(fiber.Map{"message": "Cart deleted successfully"})
}
