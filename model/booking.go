package model

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID              uuid.UUID `json:"id" from:"id"`
	Name            string    `json:"name" from:"name"`
	CheckIn         int       `json:"check_in" from:"check_in"`
	CheckOut        int       `json:"check_out" from:"check_out"`
	NumberOfPeaople int       `json:"number_of_peaople" from:"number_of_peaople"`
	TotalPrice      int       `json:"total_price" from:"total_price"`
	Status          string    `gorm:"type:enum('pending','confirmed','cancelled','expired');default:'pending'" json:"status" from:"status"`
	CreatedAt       time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt       time.Time `gorm:"type:DATETIME(0)"`
}
