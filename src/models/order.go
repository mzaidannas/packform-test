package models

import (
	"time"

	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	OrderDate  time.Time `gorm:"not null;index;index:idx_orders_date_at_brin,type:brin"`                       // Efficient range filtering/ordering of dates
	OrderName  string    `gorm:"not null;index:,type:gin,expression:order_name gin_trgm_ops" json:"OrderName"` // Full-text search on name
	CustomerId string    `gorm:"not null;index" json:"customer_id"`
	Customer   Customer  `gorm:"references:CustomerId"` // use Customer.CustomerId as references
}
