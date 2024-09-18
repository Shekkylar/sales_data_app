package models

type Order struct {
    OrderID     string `gorm:"primaryKey"`
    CustomerID  string
    DateOfSale  string
    TotalPrice  float64
    ShippingCost float64
    PaymentMethod string
}