package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllCustomers
func GetAllCustomers(c *fiber.Ctx) error {
	db := database.DB
	var customers []models.Customer
	db.Find(&customers)
	return c.JSON(fiber.Map{"status": "success", "message": "All Customers", "data": customers})
}

func ImportCustomers(c *fiber.Ctx) error {
	reader := c.Context().RequestBodyStream()
	total := repository.BulkCopy("customers", []string{"id", "name"}, &reader, 10000)
	return c.JSON(fiber.Map{"status": "success", "message": "Customers Imported", "data": total})
}
