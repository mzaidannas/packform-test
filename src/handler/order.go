package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllOrders query all orders
func GetAllOrders(c *fiber.Ctx) error {
	db := database.DB
	var orders []models.Order
	db.Find(&orders)
	return c.JSON(fiber.Map{"status": "success", "message": "All orders", "data": orders})
}
