package routes

import (
	"net/http"
	"sales_data_app/database/connection"
	"sales_data_app/services"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Not Found Error"})
	})

	// health check api endpoint for ELB
	router.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "message": ""})
	})

	router.POST("/refresh-data", func(c *gin.Context) {
		db := connection.GetDB()
		go services.LoadData(db) // Load data in the background
		c.JSON(http.StatusOK, gin.H{"message": "Data refresh initiated"})
	})

	// API to calculate total revenue for a date range
	router.GET("/revenue/total", func(c *gin.Context) {
		startDate := c.Query("start_date")
		endDate := c.Query("end_date")

		db := connection.GetDB()
		var totalRevenue float64
		db.Table("orders").
			Select("SUM(total_price) as total").
			Where("date_of_sale BETWEEN ? AND ?", startDate, endDate).
			Scan(&totalRevenue)

		c.JSON(http.StatusOK, gin.H{"total_revenue": totalRevenue})
	})

	return router
}
