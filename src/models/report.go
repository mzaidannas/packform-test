package models

import (
	"time"

	"gorm.io/gorm"
)

// Report struct
type Report struct {
	gorm.Model
	OrderName       string    `gorm:"not null;index:idx_reports_search,type:gin,expression:(order_name || ' ' || customer_name || ' ' || customer_company) gin_trgm_ops" json:"order_name"` // Full-text search on order_name
	CustomerCompany string    `gorm:"not null" json:"customer_company"`                                                                                                                     // Full-text search on company_name
	CustomerName    string    `gorm:"not null" json:"customer_name"`                                                                                                                        // Full-text search on customer_name
	OrderDate       time.Time `gorm:"not null;index:,sort:desc;index:idx_reports_order_date_brin,type:brin" json:"order_date"`                                                              // Efficient range filtering/ordering of dates
	DeliveredAmount float64   `gorm:"index" json:"delivered_amount"`
	TotalAmount     float64   `gorm:"index" json:"total_amount"`
}
