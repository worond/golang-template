package main

import (
	"app/database"
	"app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
	}))

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
