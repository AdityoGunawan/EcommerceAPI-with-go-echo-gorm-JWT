package config

import (
	"os"

	"ecommerce-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// inisialisasi database
func InitDB() {
	config := os.Getenv("CONNECTION_DB")

	var err error
	DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

// auto migrate from gorm
func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.AddressRequest{})
	DB.AutoMigrate(&models.ListOrder{})
}
