package route

import (
	"campsite_reservation/controller"
	"campsite_reservation/middleware"
	"campsite_reservation/repository"
	"campsite_reservation/usercase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := usercase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userService)

	// User
	e.POST("/api/v1/user", userController.Create)
	e.GET("/api/v1/user", userController.SelectAll, middleware.JWTMiddleware())
	e.POST("/api/v1/login", userController.Login)

	// Admin
	// e.POST("/api/v1/admin", adminController.Create, middleware.JWTMiddleware())
	// e.GET("/api/v1/admin", adminController.SelectAll, middleware.JWTMiddleware())
	// e.POST("/api/v1/admin/login", adminController.Login)

	// Campsite

}
