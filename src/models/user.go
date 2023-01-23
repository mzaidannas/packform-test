package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null;index:,type:hash" json:"password"`
	Name     string `json:"name"`
}
