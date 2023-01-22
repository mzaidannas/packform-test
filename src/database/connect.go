package database

import (
	"fmt"
	"packform-test/config"
	"packform-test/src/models"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true, // Use prepared statements to cache frequent query plans
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"pg_trgm\";")
	DB.AutoMigrate(&models.Company{}, &models.Customer{}, &models.Delivery{}, &models.Order{}, &models.OrderItem{}, &models.User{})
	fmt.Println("Database Migrated")
}
