package controller

import (
	"campsite_reservation/config"
	"campsite_reservation/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllFacilityController(c echo.Context) error {
	var facility []model.Facility

	if err := config.DB.Find(&facility).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    facility,
	})
}

func GetFacilityByIDController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var Facility model.Facility
	if err := config.DB.First(&Facility, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Facility not found")
	}

	return c.JSON(http.StatusOK, Facility)
}

func CreateFacilityController(c echo.Context) error {
	Facility := model.Facility{}

	Facility.ID = uuid.New()

	c.Bind(&Facility)

	err := config.DB.Save(&Facility).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create Facility",
		"data":    Facility,
	})
}

func UpdateFacilityController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var Facility model.Facility
	if err := config.DB.First(&Facility, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Facility not found")
	}

	if err := c.Bind(&Facility); err != nil {
		return err
	}

	if err := config.DB.Save(&Facility).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Facility)
}

func DeleteFacilityController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var Facility model.Facility
	if err := config.DB.First(&Facility, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Facility not found")
	}

	if err := config.DB.Delete(&Facility).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "Facility deleted successfully")
}
