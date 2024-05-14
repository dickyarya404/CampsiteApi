package dto

import (
	"campsite_reservation/model"
	"time"
)

type ResponseLogin struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResponseGetAll struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func FromAll(data model.User) ResponseGetAll {
	return ResponseGetAll{
		ID:        data.ID,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromAllData(data []model.User) []ResponseGetAll {
	result := []ResponseGetAll{}
	for key := range data {
		result = append(result, FromAll(data[key]))
	}
	return result
}
