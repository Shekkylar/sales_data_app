package services

import (
	"fmt"
	"log"
	"os"
	"sales_data_app/models"

	"github.com/gocarina/gocsv"
	"gorm.io/gorm"
)

func LoadData(db *gorm.DB) {
	file, err := os.Open("sales_data.csv")
	if err != nil {
		log.Fatal("Unable to read input file: ", err)
	}
	defer file.Close()

	// Parse the CSV data into SalesRecord structs
	var salesRecords []*models.SalesRecord
	if err := gocsv.UnmarshalFile(file, &salesRecords); err != nil {
		log.Fatalf("Could not parse CSV file: %v", err)
	}

	// Insert records into the database
	for _, record := range salesRecords {
		// Create Customer record
		customer := models.Customer{
			CustomerID:      record.CustomerID,
			CustomerName:    record.CustomerName,
			CustomerEmail:   record.CustomerEmail,
			CustomerAddress: record.CustomerAddress,
		}
		db.Create(&customer)

		// Create Product record
		product := models.Product{
			ProductID:   record.ProductID,
			ProductName: record.ProductName,
			Category:    record.Category,
			UnitPrice:   record.UnitPrice,
		}
		db.Create(&product)

		// Create Order record
		order := models.Order{
			OrderID:       record.OrderID,
			CustomerID:    record.CustomerID,
			DateOfSale:    record.DateOfSale,
			ShippingCost:  record.ShippingCost,
			PaymentMethod: record.PaymentMethod,
			TotalPrice:    float64(record.QuantitySold)*record.UnitPrice - record.Discount,
		}
		db.Create(&order)

		// Create OrderDetail record
		orderDetail := models.OrderDetail{
			OrderID:      record.OrderID,
			ProductID:    record.ProductID,
			QuantitySold: record.QuantitySold,
			Discount:     record.Discount,
		}
		db.Create(&orderDetail)
	}

	fmt.Println("Data loaded successfully!")

}
