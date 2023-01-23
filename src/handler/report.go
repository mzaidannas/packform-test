package handler

import (
	"packform-test/src/database"
	"packform-test/src/models"
	"packform-test/src/repository"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetReport
func GetReport(c *fiber.Ctx) error {
	db := database.DB
	var reports []models.Report
	repository.GetDatatable(db, c.Query("search"), c.Query("order")).Scan(&reports)
	return c.JSON(fiber.Map{"status": "success", "message": "All Reports", "data": reports})
}

func RefreshReports(c *fiber.Ctx) error {
	db := database.DB
	loc, _ := time.LoadLocation("Melbourne/Australia")
	start_time, _ := time.ParseInLocation(time.RFC3339, c.Query("start_time"), loc)
	end_time, _ := time.ParseInLocation(time.RFC3339, c.Query("end_time"), loc)
	count := repository.RefreshReports(start_time, end_time, db)
	return c.JSON(fiber.Map{"status": "success", "message": "Refreshed Reports", "data": count})
}
