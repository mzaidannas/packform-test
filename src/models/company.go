package models

import (
	"gorm.io/gorm"
)

// Company struct
type Company struct {
	gorm.Model
	Name string `gorm:"not null;index:,type:gin,expression:name gin_trgm_ops" json:"name"` // Full-text search on name
}
