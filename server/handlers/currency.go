package handlers

import (
	"app/database"
	"app/models"
	"context"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetCurrencies(c fiber.Ctx) error {
	ctx := context.Background()

	currencies, err := gorm.G[models.Currency](database.DB).Find(ctx)

	if err != nil {
		return c.JSON(fiber.Map{
			"data": make([]models.Currency, 0),
		})
	}

	return c.JSON(fiber.Map{
		"data": currencies,
	})
}
