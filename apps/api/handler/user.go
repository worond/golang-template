package handler

import (
	"app/database"
	"app/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {
	db := database.DB
	var user model.User
	db.First(&user, id)
	if user.Name == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

func GetUsers(c *fiber.Ctx) error {
	req := new(model.Pagination)
	if err := c.QueryParser(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid request params", "errors": err.Error()})
	}
	filter := new(model.UserFilter)
	if err := json.Unmarshal([]byte(req.Filter), &filter); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid request params", "errors": err.Error()})
	}

	db := database.DB
	var users []model.User
	var count int64

	if len(filter.Email) > 0 {
		db = db.Where("email LIKE ?", fmt.Sprintf("%%%s%%", filter.Email))
	}
	if len(filter.Name) > 0 {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
	}
	db.Order(req.Sort).Limit(req.Limit).Offset(req.Offset).Find(&users)
	db.Model(&model.User{}).Count(&count)

	return c.JSON(fiber.Map{"data": users, "total": count})
}

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User

	db.Find(&user, id)
	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"data": user})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}

	validate := validator.New()
	fmt.Println(user)
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body", "errors": err.Error()})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

// UpdateUser update user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	type UpdateUserInput struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}

	db := database.DB
	var user model.User

	db.First(&user, id)
	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	user.Name = uui.Name
	user.Email = uui.Email
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update user", "errors": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

// DeleteUser delete user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var user model.User

	db.First(&user, id)
	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
