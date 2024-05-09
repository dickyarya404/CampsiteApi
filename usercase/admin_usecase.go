package usercase

import (
	"campsite_reservation/dto"
	"campsite_reservation/model"
)

type AdminUsecase interface {
	Login(email, password string) (model.User, string, error)
	Create(payloads dto.CreateUserRequest) error
}
