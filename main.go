package main

import (
	"campsite_reservation/config"
	"campsite_reservation/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// e := echo.New()
	// cfg := config.InitConfig()
	// db := config.InitDBMysql(cfg)
	// config.InitMigrationMysql(db)

	// route.UserRoute(e, db)
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	// }))
	// e.Start(":3000")
	e := echo.New()
	// cfg := config.InitConfig()
	// db := config.InitDBMysql(cfg)
	// config.InitMigrationMysql(db)

	config.InitDB()

	// Middleware

	// Routes
	route.UserRoute(e, config.DB)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// Start server
	e.Start(":3000")

}
