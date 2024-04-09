package handler

import "github.com/gofiber/fiber/v2"

// api status
func Health(c *fiber.Ctx) error {
	var ip = c.GetRespHeader("X-Real-Ip", c.IP())
	return c.JSON(fiber.Map{"status": "ok", "ip": ip})
}
