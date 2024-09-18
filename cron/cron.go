package cron

import (
	"log"
	"sales_data_app/services"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func ScheduleDataRefresh(db *gorm.DB) {
	c := cron.New()
	c.AddFunc("@daily", func() {
		log.Println("Starting daily data refresh...")
		services.LoadData(db)
		log.Println("Data refresh completed.")
	})
	c.Start()
}
