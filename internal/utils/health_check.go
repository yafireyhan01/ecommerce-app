package utils

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "API is running smoothly",
	})
}
