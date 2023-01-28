package repository

import (
	"packform-test/src/models"
	"time"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"

	"gorm.io/gorm"
)

func GetReports(relation *gorm.DB, search string, start_date time.Time, end_date time.Time, orderCol string, orderDir string, limit int) ([]models.Report, int64, paginator.Cursor, error) {
	if search != "" {
		relation = relation.Where("(order_name || ' ' || customer_name || ' ' || customer_company) ILIKE ?", "%"+search+"%")
	}
	if !start_date.After(time.Now()) {
		relation = relation.Where("order_date >= ?", start_date)
	}
	if !end_date.Before(start_date) {
		relation = relation.Where("order_date <= ?", end_date)
	} else {
		relation = relation.Where("order_date <= CURRENT_TIMESTAMP")
	}
	var total int64
	relation = relation.Model(&models.Report{}).Count(&total)
	reports, cursor, err := GetDatatable[models.Report](relation, &search, &orderCol, &orderDir, &limit)

	// this is paginator error, e.g., invalid cursor
	if err != nil {
		return nil, 0, paginator.Cursor{}, err
	}

	return reports, total, cursor, nil
}

func RefreshReports(start_date time.Time, end_date time.Time, db *gorm.DB) int {
	sql := `
    INSERT INTO reports(order_name, customer_company, customer_name, order_date, delivered_amount, total_amount, created_at, updated_at)
    SELECT
      o.order_name,
      c.name AS customer_company,
      cus.name AS customer_name,
      o.order_date,
      SUM(d.delivered_quantity * coalesce(oi.price_per_unit, 0)) AS delivered_amount,
      SUM(coalesce(oi.quantity, 0) * coalesce(oi.price_per_unit, 0)) AS total_amount,
      NOW(),
      NOW()
    FROM orders o
    INNER JOIN customers cus ON cus.user_id = customer_id
    INNER JOIN companies c ON c.id = cus.company_id
    INNER JOIN order_items oi ON oi.order_id = o.id
    LEFT JOIN deliveries d ON d.order_item_id = oi.id
    WHERE o.order_date BETWEEN ? AND ?
    GROUP BY 1, 2, 3, 4
    ORDER BY 4 DESC;
  `
	var results []map[string]interface{}
	db.Raw(sql, start_date, end_date).Scan(results)
	return len(results)
}
