package repository

import (
	"gorm.io/gorm"
)

func GetDatatable(relation *gorm.DB, search string, order string) *gorm.DB {
	return relation.Debug().Raw(`
		SELECT
			o.order_name,
			c.name AS customer_company,
			cus.name AS customer_name,
			o.order_date,
			SUM(d.delivered_quantity * oi.price_per_unit) AS delivered_amount,
			SUM(oi.quantity * oi.price_per_unit) AS total_amount
		FROM orders o
		INNER JOIN customers cus ON cus.user_id = customer_id
		INNER JOIN companies c ON c.id = cus.company_id
		INNER JOIN order_items oi ON oi.order_id = o.id
		INNER JOIN deliveries d ON d.order_item_id = oi.id
		WHERE o.order_name % ? OR c.name % ? OR cus.name % ?
		GROUP BY 1, 2, 3, 4
		ORDER BY 4 DESC;
	`, search, search, search)
}
