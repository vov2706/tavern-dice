package main

import (
	"app/config"
	"app/routes"
	"log"

	"app/database"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "Tavern Dice",
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: false,
	}))

	database.Connect()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.Config("APP_PORT")))
}
