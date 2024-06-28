package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
)

func setupRoutes(app *fiber.App) {
	app.Get("/health-check", utils.HealthCheck)

}
