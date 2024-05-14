package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID              uuid.UUID      `json:"id" from:"id"`
	CampsiteID      uuid.UUID      `json:"campsite_id" from:"campsite_id"`
	UserID          uuid.UUID      `json:"user_id" from:"user_id"`
	Name            string         `json:"name" from:"name"`
	CheckIn         string         `json:"check_in" from:"check_in"`
	CheckOut        string         `json:"check_out" from:"check_out"`
	NumberOfPeaople int            `json:"number_of_peaople" from:"number_of_peaople"`
	TotalPrice      int            `json:"total_price" from:"total_price"`
	Status          string         `gorm:"type:enum('pending','confirmed','cancelled','expired');default:'pending'" json:"status" from:"status"`
	CreatedAt       time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt       time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt       gorm.DeletedAt `gorm:"type:DATETIME(0)"`
}

type BookingDetail struct {
	ID   uuid.UUID `json:"id" from:"id"`
	Name string    `json:"name" from:"name"`
}
