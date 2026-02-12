package routes

import (
	"app/handlers"
	"app/middlewares"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Auth
	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.Register)

	// Profile
	api.Get("/profile", middlewares.Protected(), handlers.GetProfile)

	// 404
	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound) // => 404 "Not Found"
	})
}
