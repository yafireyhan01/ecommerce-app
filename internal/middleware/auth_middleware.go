package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
)

func AuthorizeJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header required"})
		}
		tokenString := strings.TrimPrefix(authHeader, BEARER_SCHEMA)

		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			case jwt.ValidationErrorExpired:
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized, token expired"})
			default:
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}
		}

		c.Locals("UserID", claims.Guid)
		c.Locals("UserEmail", claims.Email)
		c.Locals("UserRole", claims.Role)

		return c.Next()
	}
}

func AuthorizeUserRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("UserRole").(string)
		if role != userRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}
		return c.Next()
	}
}
