package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected(c *fiber.Ctx) error {
	token := c.Cookies("Authorization", "")
	if len(token) == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
