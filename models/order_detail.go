package models

type OrderDetail struct {
    OrderID     string `gorm:"primaryKey"`
    ProductID   string `gorm:"primaryKey"`
    QuantitySold int
    Discount     float64
}