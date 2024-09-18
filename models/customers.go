package models

type Customer struct {
    CustomerID   string `gorm:"primaryKey"`
    CustomerName string
    CustomerEmail string
    CustomerAddress string
}