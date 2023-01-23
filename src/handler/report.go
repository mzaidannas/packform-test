package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"

	"github.com/gofiber/fiber/v2"
)

// GetReport
func GetReport(c *fiber.Ctx) error {
	db := database.DB
	var reports []models.Report
	repository.GetDatatable(db, c.Query("search"), c.Query("order")).Scan(&reports)
	return c.JSON(fiber.Map{"status": "success", "message": "All Customers", "data": reports})
}
