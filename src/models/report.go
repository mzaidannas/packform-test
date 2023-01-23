package models

import "time"

// Company struct
type Report struct {
	OrderName       string    `json:"order_name"`
	CustomerName    string    `json:"customer_name"`
	OrderDate       time.Time `json:"order_date"`
	DeliveredAmount float64   `json:"delivered_amount"`
	TotalAmount     float64   `json:"total_amount"`
}
