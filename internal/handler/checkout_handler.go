package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/service"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
)

type CheckoutHandler struct {
	checkoutService service.CheckoutService
	validate        *validator.Validate
}

func NewCheckoutHandler(checkoutService service.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{
		checkoutService: checkoutService,
		validate:        validator.New(),
	}
}

func (h *CheckoutHandler) CreateCheckout(c *fiber.Ctx) error {
	userGuid := c.Locals("UserGuid").(string)
	var input struct {
		CartGuid string `json:"cart_guid" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	cartGuid := uuid.FromStringOrNil(input.CartGuid)
	totalPrice, err := h.checkoutService.CalculateTotalPrice(cartGuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not calculate total price"})
	}

	checkout := &models.Checkout{
		Guid:       utils.NewUUID(),
		UserGuid:   uuid.FromStringOrNil(userGuid),
		CartGuid:   cartGuid,
		TotalPrice: totalPrice,
		Status:     "PENDING",
	}

	if err := h.checkoutService.CreateCheckout(checkout); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create checkout"})
	}

	return c.JSON(fiber.Map{
		"user_guid":   checkout.UserGuid,
		"cart_guid":   checkout.CartGuid,
		"total_price": checkout.TotalPrice,
		"status":      checkout.Status,
	})
}
