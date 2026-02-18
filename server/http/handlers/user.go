package handlers

import (
	"app/database"
	"app/http/responses"
	"app/models"
	"context"
	"errors"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetProfile(c fiber.Ctx) error {
	user, err := GetAuthUser(c)

	if err != nil {
		return err
	}

	return c.JSON(responses.UserResponse{
		Data: responses.NewUserResource(*user),
	})
}

func GetAuthUser(c fiber.Ctx) (*models.User, error) {
	token := jwtware.FromContext(c)

	if token == nil {
		return nil, fiber.ErrUnauthorized
	}

	claims := token.Claims.(jwt.MapClaims)
	userId, ok := claims["sub"].(float64)

	if !ok {
		return nil, fiber.ErrUnauthorized
	}

	user, err := getUserByID(uint(userId))

	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	return &user, nil
}

func getUserByID(id uint) (models.User, error) {
	ctx := context.Background()
	user, err := gorm.G[models.User](database.DB).
		Where("id = ?", id).
		Preload("Balances", func(db gorm.PreloadBuilder) error {
			subQuery := database.DB.Table("currencies").Select("id").Where("slug = ?", models.GOLD)

			db.Where("currency_id = (?)", subQuery)

			return nil
		}).
		Preload("Balances.Currency", nil).
		First(ctx)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(username, password string) (models.User, error) {
	ctx := context.Background()
	currency, err := gorm.G[models.Currency](database.DB).Where("slug = ?", models.BRONZE).First(ctx)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: username,
		Password: password,
		Balances: []models.Balance{
			{CurrencyID: currency.ID, Amount: 1000},
		},
	}

	err = gorm.G[models.User](database.DB).Create(ctx, &user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func getUserByUsername(username string) (*models.User, error) {
	ctx := context.Background()

	user, err := gorm.G[models.User](database.DB).Where("username = ?", username).First(ctx)

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
