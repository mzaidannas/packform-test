package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllCustomers
func GetAllCustomers(c *fiber.Ctx) error {
	db := database.DB
	var customers []models.Customer
	db.Find(&customers)
	return c.JSON(fiber.Map{"status": "success", "message": "All Customers", "data": customers})
}
