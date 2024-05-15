package controller

import (
	"campsite_reservation/config"
	"campsite_reservation/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllAvailabilityController(c echo.Context) error {
	var availability []model.Availability

	if err := config.DB.Find(&availability).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    availability,
	})
}

func GetAvailabilityByIDController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var availability model.Availability
	if err := config.DB.First(&availability, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Availability not found")
	}

	return c.JSON(http.StatusOK, availability)
}

func CreateAvailabilityController(c echo.Context) error {
	availability := model.Availability{}

	availability.ID = uuid.New()

	c.Bind(&availability)

	err := config.DB.Save(&availability).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create Availability",
		"data":    availability,
	})
}

func UpdateAvailabilityController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var availability model.Availability
	if err := config.DB.First(&availability, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Availability not found")
	}

	if err := c.Bind(&availability); err != nil {
		return err
	}

	if err := config.DB.Save(&availability).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, availability)
}

func DeleteAvailabilityController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var availability model.Availability
	if err := config.DB.First(&availability, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Availability not found")
	}

	if err := config.DB.Delete(&availability).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "Availability deleted successfully")
}
