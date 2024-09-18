package bootstrap

import (
	"sales_data_app/cron"
	"sales_data_app/database/connection"
	"sales_data_app/routes"
)

func Init() {
	connection.PostgresqlConnection()
	cron.ScheduleDataRefresh(connection.GetDB())
	routes.Init()
}
