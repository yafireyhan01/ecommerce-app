package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
)

type PaymentService interface {
	PayCheckout(userGuid string, checkoutGuid uuid.UUID) error
}

type paymentService struct {
	paymentRepo  repository.PaymentRepository
	checkoutRepo repository.CheckoutRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository, checkoutRepo repository.CheckoutRepository) PaymentService {
	return &paymentService{
		paymentRepo:  paymentRepo,
		checkoutRepo: checkoutRepo,
	}
}

func (s *paymentService) PayCheckout(userGuid string, checkoutGuid uuid.UUID) error {
	checkout, err := s.checkoutRepo.FindByID(checkoutGuid.String())
	if err != nil {
		return err
	}

	if checkout.UserGuid.String() != userGuid || checkout.Status != "PENDING" {
		return fiber.NewError(fiber.StatusForbidden, "Unauthorized or invalid checkout status")
	}

	checkout.Status = "PAID"
	if err := s.checkoutRepo.Update(checkout); err != nil {
		return err
	}

	payment := &models.Payment{
		Guid:         utils.NewUUID(),
		CheckoutGuid: checkout.Guid,
	}
	return s.paymentRepo.Create(payment)
}
