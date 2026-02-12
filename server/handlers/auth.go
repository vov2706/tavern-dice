package handlers

import (
	"app/config"
	"app/models"
	"app/utils"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

func Login(c fiber.Ctx) error {
	userData, err := validateUserData(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := getUserByUsername(userData.Username)

	const dummyHash = "2a$10$7zFqzDbD3RrlkMTczbXG9OWZ0FLOXjIxXzSZ.QZxkVXjXcx7QZQiC"

	if err != nil {
		// prevents timing attacks
		utils.IsValidPassword(dummyHash, userData.Password)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	if !utils.IsValidPassword(user.Password, userData.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	token, err := createToken(*user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logged in",
		"token":   token,
	})
}

func Register(c fiber.Ctx) error {
	userData, err := validateUserData(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if isUserExists(userData.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	hashedPassword := utils.GeneratePassword(userData.Password)
	user, err := CreateUser(userData.Username, hashedPassword)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := createToken(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success registration",
		"token":   token,
	})
}

func isUserExists(username string) bool {
	_, err := getUserByUsername(username)

	return err == nil
}

func validateUserData(c fiber.Ctx) (*LoginInput, error) {
	input := new(LoginInput)

	if err := c.Bind().Body(input); err != nil {
		return nil, err
	}

	// Validation logic
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return nil, e
		}
	}

	return input, nil
}

func createToken(u models.User) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err = token.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return t, nil
}
