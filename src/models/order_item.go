package models

import (
	"gorm.io/gorm"
)

// OrderItem struct
type OrderItem struct {
	gorm.Model
	OrderId      uint `gorm:"not null;index" json:"order_id"`
	Order        Order
	PricePerUnit float32 `gorm:"null;index:,sort:desc,null:last" json:"price_per_unit"`
	Quantity     uint    `gorm:"not null;index:,sort:desc,null:last" json:"quantity"`
	Product      string  `gorm:"not null"`
}
