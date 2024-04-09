package handler

import (
	"app/config"
	"app/repository"
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `validate:"required,email" json:"email"`
		Password string `validate:"required" json:"password"`
	}
	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request body", "errors": err.Error()})
	}

	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal Server Error", "data": err})
	} else if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": err})
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.GetInt("JWT_EXPIRES"))).Unix()

	t, err := token.SignedString([]byte(config.GetString("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func GetToken(c *fiber.Ctx) error {
	type ExchangeInput struct {
		Code         string `json:"code"`
		CodeVerifier string `json:"code_verifier"`
	}

	input := new(ExchangeInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
	}
	fmt.Println(input)
	oauth2Conf := config.OAuth2Config

	token, err := oauth2Conf.Exchange(context.Background(), input.Code, oauth2.VerifierOption(input.CodeVerifier))
	if err != nil {
		fmt.Println(err)
		return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Code-Token Exchange Failed", "token": nil})
	}

	return c.JSON(token)
}
