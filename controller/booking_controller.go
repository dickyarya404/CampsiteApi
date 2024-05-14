package controller

import (
	"campsite_reservation/config"
	"campsite_reservation/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllBookingController(c echo.Context) error {
	var booking []model.Booking

	if err := config.DB.Find(&booking).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    booking,
	})
}

// Get package by ID
func GetBookingByIDController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var booking model.Booking
	if err := config.DB.First(&booking, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "booking not found")
	}

	return c.JSON(http.StatusOK, booking)
}

// Add new package
func CreateBookingController(c echo.Context) error {
	booking := model.Booking{}

	booking.ID = uuid.New()

	c.Bind(&booking)

	err := config.DB.Save(&booking).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create booking",
		"data":    booking,
	})
}

// Update package
func UpdateBookingController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var booking model.Booking
	if err := config.DB.First(&booking, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "booking not found")
	}

	if err := c.Bind(&booking); err != nil {
		return err
	}

	if err := config.DB.Save(&booking).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, booking)
}

func DeleteBookingController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var booking model.Booking
	if err := config.DB.First(&booking, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "booking not found")
	}

	if err := config.DB.Delete(&booking).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "booking deleted successfully")
}
