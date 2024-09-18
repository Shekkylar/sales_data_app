package connection

import (
	"log"
	"sales_data_app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func PostgresqlConnection() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=sales_data port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	d.AutoMigrate(&models.Order{}, &models.Customer{}, &models.Product{}, &models.OrderDetail{})
	db = d
}

func GetDB() *gorm.DB {
	return db
}
