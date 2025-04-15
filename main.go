package main

import (
	"fiber-auth-app/config"
	"fiber-auth-app/middleware"
	"fiber-auth-app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to DB
	config.ConnectDatabase()

	// Create Fiber app
	app := fiber.New()

	// Panic Handler
	app.Use(middleware.CustomRecoverPanic())

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
