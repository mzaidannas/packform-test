package models

import (
	"time"

	"gorm.io/gorm"
)

var loc time.Location

// Order struct
type Order struct {
	gorm.Model
	OrderDate  time.Time `gorm:"not null;index;index:idx_orders_date_at_brin,type:brin"`                       // Efficient range filtering/ordering of dates
	OrderName  string    `gorm:"not null;index:,type:gin,expression:order_name gin_trgm_ops" json:"OrderName"` // Full-text search on name
	CustomerId string    `gorm:"not null;index" json:"customer_id"`
	Customer   Customer  `gorm:"references:UserId;foreignKey:CustomerId"` // use Customer.CustomerId as references
}

func OrderCsv(row *[]string) []interface{} {
	loc, _ := time.LoadLocation("Europe/Berlin")
	new := make([]interface{}, len(*row))
	for k, v := range *row {
		if k == 1 {
			new[k], _ = time.ParseInLocation(time.RFC3339, v, loc)
		} else {
			new[k] = v
		}
	}
	return new
}
