package main

import (
	"campsite_reservation/config"
	"campsite_reservation/route"

	// "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// _ "campsite_reservation/docs" // Import the generated docs
	// _ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
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
