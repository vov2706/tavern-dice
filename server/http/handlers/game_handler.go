package handlers

import (
	"app/http/inputs"
	"app/http/responses"
	"app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type GameHandler struct {
	gameService *services.GameService
}

func NewGameHandler(gameService *services.GameService) *GameHandler {
	return &GameHandler{gameService: gameService}
}

func (handler *GameHandler) CreateGame(c fiber.Ctx) error {
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
	if err := input.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	authUser, err := GetAuthUser(c)

	if err != nil {
		return err
	}

	game, err := handler.gameService.CreateGame(authUser, input)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(responses.NewGameResource(*game))
}
