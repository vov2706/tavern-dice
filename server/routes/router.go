package routes

import (
	"app/database"
	"app/http/handlers"
	"app/http/middlewares"
	"app/repositories"
	"app/services"

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

	// Games
	balanceRepo := repositories.NewBalanceRepository(database.DB)
	gameRepo := repositories.NewGameRepository(database.DB)
	currencyRepo := repositories.NewCurrencyRepository(database.DB)
	gameService := services.NewGameService(balanceRepo, currencyRepo, gameRepo)
	gameHandler := handlers.NewGameHandler(gameService)
	api.Post("/games", middlewares.Protected(), gameHandler.CreateGame)

	// Currencies
	api.Get("/currencies", handlers.GetCurrencies)

	// 404
	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound) // => 404 "Not Found"
	})
}
