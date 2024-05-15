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
	e.POST("/api/v1/register", userController.Create)
	e.GET("/api/v1/user", userController.SelectAll, middleware.JWTMiddleware())
	e.POST("/api/v1/login", userController.Login)

	// Campsite

	e.GET("/api/v1/campsite", controller.GetAllCampsiteController)
	e.GET("/api/v1/campsite/:id", controller.GetCampsiteByIDController)
	e.POST("/api/v1/campsite", controller.CreateCampsiteController)
	e.PUT("/api/v1/campsite/:id", controller.UpdateCampsiteController)
	e.DELETE("/api/v1/campsite/:id", controller.DeleteCampsiteController)

	// Booking
	e.GET("/api/v1/booking", controller.GetAllBookingController)
	e.GET("/api/v1/booking/:id", controller.GetBookingByIDController)
	e.POST("/api/v1/booking", controller.CreateBookingController)
	e.PUT("/api/v1/booking/:id", controller.UpdateBookingController)
	e.DELETE("/api/v1/booking/:id", controller.DeleteBookingController)

	// Facility
	e.GET("/api/v1/facility", controller.GetAllFacilityController)
	e.GET("/api/v1/facility/:id", controller.GetFacilityByIDController)
	e.POST("/api/v1/facility", controller.CreateFacilityController)
	e.PUT("/api/v1/facility/:id", controller.UpdateFacilityController)
	e.DELETE("/api/v1/facility/:id", controller.DeleteFacilityController)

	// Review
	e.GET("/api/v1/review", controller.GetAllReviewController)
	e.GET("/api/v1/review/:id", controller.GetReviewByIDController)
	e.POST("/api/v1/review", controller.CreateReviewController)
	e.PUT("/api/v1/review/:id", controller.UpdateReviewController)
	e.DELETE("/api/v1/review/:id", controller.DeleteReviewController)

	// Availability
	e.GET("/api/v1/availability", controller.GetAllAvailabilityController)
	e.GET("/api/v1/availability/:id", controller.GetAvailabilityByIDController)
	e.POST("/api/v1/availability", controller.CreateAvailabilityController)
	e.PUT("/api/v1/availability/:id", controller.UpdateAvailabilityController)
	e.DELETE("/api/v1/availability/:id", controller.DeleteAvailabilityController)

}
