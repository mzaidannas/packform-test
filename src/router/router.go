package router

import (
	"packform-test/src/handler"
	"packform-test/src/middlewares"

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

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middlewares.Protected(), handler.UpdateUser)
	user.Delete("/:id", middlewares.Protected(), handler.DeleteUser)

	// Order
	order := api.Group("/order")
	order.Get("/", handler.GetAllOrders)
}
