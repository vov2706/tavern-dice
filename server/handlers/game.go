package handlers

import (
	"app/database"
	"app/inputs"
	"app/models"
	"app/responses"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateGame(c fiber.Ctx) error {
	input := new(inputs.CreateGameInput)

	if err := c.Bind().Body(input); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"field":   e.Field(),
					"message": e.Error(),
				})
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// custom validation
	err := input.Validate()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	authUser, err := GetAuthUser(c)

	if err != nil {
		return err
	}

	ctx := context.Background()
	currency, err := gorm.G[models.Currency](database.DB).First(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get currency",
		})
	}

	game := models.Game{
		Code:          uuid.New().String(),
		CreatorID:     authUser.ID,
		CurrencyID:    input.CurrencyID,
		Bet:           input.Bet,
		WinningPoints: input.WinningPoints,
		JoinType:      input.JoinType,
	}
	err = gorm.G[models.Game](database.DB).Create(ctx, &game)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create game",
		})
	}

	game.Currency = currency

	return c.JSON(fiber.Map{
		"data": responses.NewGameResource(game),
	})
}
