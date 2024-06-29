package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yafireyhan01/synapsis-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Run() error {
	app := fiber.New()

	// Load config
	config.LoadConfig()

	// Connect to database
	dsn := "host=" + config.AppConfig.DBHost +
		" user=" + config.AppConfig.DBUser +
		" password=" + config.AppConfig.DBPassword +
		" dbname=" + config.AppConfig.DBName +
		" port=" + config.AppConfig.DBPort +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Setup routes
	setupRoutes(app, db)

	// Start server
	port := config.AppConfig.APIPort
	if port == "" {
		port = "8080"
	}
	return app.Listen(":" + port)
}
