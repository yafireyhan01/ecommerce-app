package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/internal/handler"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
	"github.com/yafireyhan01/synapsis-test/internal/service"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/health-check", utils.HealthCheck)

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	app.Post("/register", authHandler.Register)
}
