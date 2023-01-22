package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllDeliveries
func GetAllDeliveries(c *fiber.Ctx) error {
	db := database.DB
	var deliveries []models.Delivery
	db.Find(&deliveries)
	return c.JSON(fiber.Map{"status": "success", "message": "All Deliveries", "data": deliveries})
}

func ImportDeliveries(c *fiber.Ctx) error {
	reader := c.Context().RequestBodyStream()
	total := repository.BulkCopy("deliveries", []string{"id", "order_item_id", "delivered_quantity"}, &reader, 10000)
	return c.JSON(fiber.Map{"status": "success", "message": "Deliveries Imported", "data": total})
}
