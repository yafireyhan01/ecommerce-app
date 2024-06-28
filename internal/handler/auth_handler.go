package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/internal/models"
	"github.com/yafireyhan01/synapsis-test/internal/service"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	authService service.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		Role     string `json:"role" validate:"required,oneof=CUSTOMER SELLER"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := h.validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//existingUser, err := h.authService.FindByEmail(input.Email)
	//if err == nil && existingUser != nil {
	//	return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already in use"})
	//}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	user := &models.User{
		Guid:     utils.NewUUID(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := h.authService.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
