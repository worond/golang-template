package router

import (
	"app/handler"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/token", handler.GetToken)

	// User
	user := api.Group("/users")
	user.Get("/", handler.GetUsers)         //getList
	user.Get("/:id", handler.GetUser)       //getOne
	user.Post("/", handler.CreateUser)      //create
	user.Patch("/:id", handler.UpdateUser)  //update
	user.Delete("/:id", handler.DeleteUser) //delete

	// Product
	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
