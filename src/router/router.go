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
	auth.Post("/register", handler.Register)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Patch("/:id", middlewares.Protected(), handler.UpdateUser)
	user.Delete("/:id", middlewares.Protected(), handler.DeleteUser)

	// Order
	order := api.Group("/order")
	order.Get("/", middlewares.Protected(), handler.GetAllOrders)
	order.Post("/import", middlewares.Protected(), handler.ImportOrders)

	// Customer
	customer := api.Group("/customer")
	customer.Get("/", middlewares.Protected(), handler.GetAllCustomers)
	customer.Post("/import", middlewares.Protected(), handler.ImportCustomers)

	// Company
	company := api.Group("/company")
	company.Get("/", middlewares.Protected(), handler.GetAllCompanies)
	company.Post("/import", middlewares.Protected(), handler.ImportCompanies)

	// OrderItem
	orderItem := api.Group("/order-item")
	orderItem.Get("/", middlewares.Protected(), handler.GetAllOrderItems)
	orderItem.Post("/import", middlewares.Protected(), handler.ImportOrderItems)

	// Delivery
	delivery := api.Group("/delivery")
	delivery.Get("/", middlewares.Protected(), handler.GetAllDeliveries)
	delivery.Post("/import", middlewares.Protected(), handler.ImportDeliveries)

	// Report
	report := api.Group("/report")
	report.Get("/", middlewares.Protected(), handler.GetReport)
	report.Get("/refresh", middlewares.Protected(), handler.RefreshReports)
}
