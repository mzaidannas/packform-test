package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllCompanies
func GetAllCompanies(c *fiber.Ctx) error {
	db := database.DB
	var companies []models.Company
	db.Find(&companies)
	return c.JSON(fiber.Map{"status": "success", "message": "All Companies", "data": companies})
}

func ImportCompanies(c *fiber.Ctx) error {
	reader := c.Context().RequestBodyStream()
	total := repository.BulkCopy("companies", []string{"id", "name"}, &reader, 10000)
	return c.JSON(fiber.Map{"status": "success", "message": "Companies Imported", "data": total})
}
