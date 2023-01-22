package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllOrderItems
func GetAllOrderItems(c *fiber.Ctx) error {
	db := database.DB
	var orderItems []models.OrderItem
	db.Find(&orderItems)
	return c.JSON(fiber.Map{"status": "success", "message": "All OrderItems", "data": orderItems})
}

func ImportOrderItems(c *fiber.Ctx) error {
	reader := c.Context().RequestBodyStream()
	total := repository.BulkCopy("orderItems", []string{"id", "order_id", "price_per_unit", "quantity", "product"}, &reader, 10000)
	return c.JSON(fiber.Map{"status": "success", "message": "OrderItems Imported", "data": total})
}
