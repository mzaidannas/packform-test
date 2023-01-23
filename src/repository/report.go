package repository

import (
	"time"

	"gorm.io/gorm"
)

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
	db.Debug().Raw(sql, start_date, end_date).Scan(results)
	return len(results)
}
