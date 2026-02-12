package handlers

import (
	"app/database"
	"app/models"
	"app/responses"
	"context"
	"errors"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetProfile(c fiber.Ctx) error {
	token := jwtware.FromContext(c)

	if token == nil {
		return fiber.ErrUnauthorized
	}

	claims := token.Claims.(jwt.MapClaims)
	userId, ok := claims["sub"].(float64)

	if !ok {
		return fiber.ErrUnauthorized
	}

	user, err := GetUser(uint(userId))

	if err != nil {
		return fiber.ErrUnauthorized
	}

	return c.JSON(responses.UserResponse{
		Data: responses.NewUserResource(user),
	})
}

func GetUser(id uint) (models.User, error) {
	ctx := context.Background()
	user, err := gorm.G[models.User](database.DB).Where("id = ?", id).First(ctx)

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(username, password string) (models.User, error) {
	ctx := context.Background()

	user := models.User{
		Username: username,
		Password: password,
	}

	err := gorm.G[models.User](database.DB).Create(ctx, &user)

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
