package controller

import (
	"campsite_reservation/config"
	"campsite_reservation/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllReviewController(c echo.Context) error {
	var review []model.Review

	if err := config.DB.Find(&review).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    review,
	})
}

func GetReviewByIDController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var review model.Review
	if err := config.DB.First(&review, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Review not found")
	}

	return c.JSON(http.StatusOK, review)
}

func CreateReviewController(c echo.Context) error {
	Review := model.Review{}

	Review.ID = uuid.New()

	c.Bind(&Review)

	err := config.DB.Save(&Review).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create Review",
		"data":    Review,
	})
}

func UpdateReviewController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var Review model.Review
	if err := config.DB.First(&Review, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Review not found")
	}

	if err := c.Bind(&Review); err != nil {
		return err
	}

	if err := config.DB.Save(&Review).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Review)
}

func DeleteReviewController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	var Review model.Review
	if err := config.DB.First(&Review, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Review not found")
	}

	if err := config.DB.Delete(&Review).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, "Review deleted successfully")
}
