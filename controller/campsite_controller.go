package controller

import (
	"campsite_reservation/config"
	"campsite_reservation/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllCampsiteController(c echo.Context) error {
	var campsite []model.Campsite

	if err := config.DB.Find(&campsite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    campsite,
	})
}

// Get package by ID
func GetCampsiteByIDController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var campsite model.Campsite
	if err := config.DB.First(&campsite, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Campsite not found")
	}

	return c.JSON(http.StatusOK, campsite)
}

// Add new package
func CreateCampsiteController(c echo.Context) error {
	campsite := model.Campsite{}

	campsite.ID = uuid.New()

	c.Bind(&campsite)

	err := config.DB.Save(&campsite).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create campsite",
		"data":    campsite,
	})
}

// Update package
func UpdateCampsiteController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var campsite model.Campsite
	if err := config.DB.First(&campsite, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Campsite not found")
	}

	if err := c.Bind(&campsite); err != nil {
		return err
	}

	if err := config.DB.Save(&campsite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, campsite)
}

func DeleteCampsiteController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var campsite model.Campsite
	if err := config.DB.First(&campsite, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Campsite not found")
	}

	if err := config.DB.Delete(&campsite).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "Campsite deleted successfully")
}
