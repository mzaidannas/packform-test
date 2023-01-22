package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllOrders query all orders
func GetAllOrders(c *fiber.Ctx) error {
	db := database.DB
	var orders []models.Order
	db.Find(&orders)
	return c.JSON(fiber.Map{"status": "success", "message": "All orders", "data": orders})
}

func ImportOrders(c *fiber.Ctx) error {
	reader := c.Context().RequestBodyStream()
	total := repository.BulkCopy("orders", []string{"id", "created_at", "order_name", "customer_id"}, &reader, 10000)
	return c.JSON(fiber.Map{"status": "success", "message": "Companies Imported", "data": total})
}
