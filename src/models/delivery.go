package models

import (
	"gorm.io/gorm"
)

// Delivery struct
type Delivery struct {
	gorm.Model
	OrderItemId       uint `gorm:"not null;index" json:"order_item_id"`
	OrderItem         OrderItem
	DeliveredQuantity uint `gorm:"not null;index:,sort:desc,null:last" json:"delivered_quantity"`
}
