package handler

import (
	"packform-test/src/database"
	"packform-test/src/repository"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetReport
func GetReport(c *fiber.Ctx) error {
	db := database.DB
	loc, _ := time.LoadLocation("Melbourne/Australia")
	start_date, _ := time.ParseInLocation(time.RFC3339, c.Query("start_date"), loc)
	end_date, _ := time.ParseInLocation(time.RFC3339, c.Query("end_date"), loc)
	limit, _ := strconv.Atoi(c.Query("limit"))
	reports, total, cursor, err := repository.GetReports(db, c.Query("search"), start_date, end_date, c.Query("order_col"), c.Query("order"), limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't retrieve reports", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All Reports", "data": reports, "cursor": cursor, "total": total})
}

func RefreshReports(c *fiber.Ctx) error {
	db := database.DB
	loc, _ := time.LoadLocation("Melbourne/Australia")
	start_time, _ := time.ParseInLocation(time.RFC3339, c.Query("start_time"), loc)
	end_time, _ := time.ParseInLocation(time.RFC3339, c.Query("end_time"), loc)
	count := repository.RefreshReports(start_time, end_time, db)
	return c.JSON(fiber.Map{"status": "success", "message": "Refreshed Reports", "data": count})
}
