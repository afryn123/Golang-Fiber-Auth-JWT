package routes

import (
	"fiber-auth-app/handlers"
	"fiber-auth-app/middleware"
	"fiber-auth-app/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userRepo := repository.NewUserRepository()
	authHandler := handlers.NewAuthHandler(userRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	api := app.Group("/api")

	api.Get("/healthCheck", handlers.HealthCheck)

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	api.Get("/profile", middleware.JWTProtected(), userHandler.Profile)
}
