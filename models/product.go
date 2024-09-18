package models

type Product struct {
    ProductID    string `gorm:"primaryKey"`
    ProductName  string
    Category     string
    UnitPrice    float64
}