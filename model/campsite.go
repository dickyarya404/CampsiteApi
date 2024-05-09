package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campsite struct {
	*gorm.Model
	ID        uuid.UUID `json:"campsite_id" from:"campsite_id" grom:"campsite_id"`
	Name      string    `json:"name" from:"name"`
	Location  string    `json:"location" from:"location"`
	Latitude  float64   `json:"latitude" from:"latitude"`
	Lontitude float64   `json:"lontitude" from:"lontitude"`
	Luas      int       `json:"luas" from:"luas"`
	Price     int       `json:"price" from:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
