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

	// User
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	// Category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Public routes
	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)

	// Protected routes
	api := app.Group("/api")
	api.Use(middleware.AuthorizeJWT())
	api.Get("/protected-route", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "You are authorized!"})
	})

	// Seller for Seller role
	sellerGroup := api.Group("/seller")
	sellerGroup.Use(middleware.AuthorizeUserRole("SELLER"))

	// Category for Seller role
	sellerGroup.Post("/categories", categoryHandler.CreateCategory)
	sellerGroup.Put("/categories/:guid", categoryHandler.UpdateCategory)
	sellerGroup.Delete("/categories/:guid", categoryHandler.DeleteCategory)
	sellerGroup.Get("/categories/:guid", categoryHandler.GetCategoryByID)
	sellerGroup.Get("/categories", categoryHandler.GetAllCategories)

	// Product for Seller role
	sellerGroup.Post("/products", productHandler.CreateProduct)
	sellerGroup.Put("/products/:guid", productHandler.UpdateProduct)
	sellerGroup.Delete("/products/:guid", productHandler.DeleteProduct)
	sellerGroup.Get("/products/:guid", productHandler.GetProductByID)
	sellerGroup.Get("/products", productHandler.GetAllProducts)
	sellerGroup.Get("/categories/:category_guid/products", productHandler.GetProductsByCategory)

	// Customer
	customerGroup := api.Group("/customer")
	customerGroup.Use(middleware.AuthorizeUserRole("CUSTOMER"))

	// Product for Customer role
	customerGroup.Get("/products/:guid", productHandler.GetProductByID)
	customerGroup.Get("/products", productHandler.GetAllProducts)
	customerGroup.Get("/categories/:category_guid/products", productHandler.GetProductsByCategory)
}
