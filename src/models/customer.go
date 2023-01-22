package models

import (
	"gorm.io/gorm"
)

// Customer struct
type Customer struct {
	gorm.Model
	UserId      string   `gorm:"uniqueIndex;not null" json:"user_id"`
	Login       string   `gorm:"uniqueIndex;not null" json:"login"`
	Password    string   `gorm:"not null" json:"password"`
	Name        string   `gorm:"not null;index:,type:gin,expression:name gin_trgm_ops" json:"name"` // Full-text search on name
	CompanyId   int      `gorm:"not null;index" json:"company_id"`
	CreditCards []string `gorm:"type:text[]" json:"credit_cards"`
	Company     Company
}
