package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/internal/handler"
	"github.com/yafireyhan01/synapsis-test/internal/middleware"
	"github.com/yafireyhan01/synapsis-test/internal/repository"
	"github.com/yafireyhan01/synapsis-test/internal/service"
	"github.com/yafireyhan01/synapsis-test/internal/utils"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	// Health check route
	app.Get("/health-check", utils.HealthCheck)

	// Initialize repositories, services, and handlers
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	// Public routes
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	// Protected routes
	api := app.Group("/api")
	api.Use(middleware.AuthorizeJWT())
	api.Get("/protected-route", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "You are authorized!"})
	})

	// Role-specific routes
	sellerGroup := api.Group("/seller")
	sellerGroup.Use(middleware.AuthorizeUserRole("SELLER"))
	sellerGroup.Post("/create-product", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Seller authorized to create product"})
	})

	customerGroup := api.Group("/customer")
	customerGroup.Use(middleware.AuthorizeUserRole("CUSTOMER"))
	customerGroup.Post("/make-order", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Customer authorized to make order"})
	})
}
