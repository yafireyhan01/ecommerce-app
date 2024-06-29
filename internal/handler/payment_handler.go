package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/yafireyhan01/synapsis-test/internal/service"
)

type PaymentHandler struct {
	paymentService service.PaymentService
	validate       *validator.Validate
}

func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		validate:       validator.New(),
	}
}

func (h *PaymentHandler) PayCheckout(c *fiber.Ctx) error {
	userGuid := c.Locals("UserGuid").(string)
	var input struct {
		CheckoutGuid string `json:"checkout_guid" validate:"required"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	checkoutGuid := uuid.FromStringOrNil(input.CheckoutGuid)

	if err := h.paymentService.PayCheckout(userGuid, checkoutGuid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not process payment"})
	}

	return c.JSON(fiber.Map{"message": "Payment successful"})
}
