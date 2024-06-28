package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/internal/config"
)

func Run() error {
	app := fiber.New()

	// Load config
	config.LoadConfig()

	// Setup routes
	setupRoutes(app)

	// Start server
	return app.Listen(":8080")
}
