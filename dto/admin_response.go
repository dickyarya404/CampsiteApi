package dto

import (
	"campsite_reservation/model"
	"time"
)

type ResponseLoginAdmin struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResponseGetAllAdmin struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func FromAllAdmin(data model.Admin) ResponseGetAllAdmin {
	return ResponseGetAllAdmin{
		ID:        data.ID,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromAllDataAdmin(data []model.Admin) []ResponseGetAllAdmin {
	result := []ResponseGetAllAdmin{}
	for key := range data {
		result = append(result, FromAllAdmin(data[key]))
	}
	return result
}
