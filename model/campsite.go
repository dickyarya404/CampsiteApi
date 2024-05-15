package model

import (
	"time"

	"github.com/google/uuid"
)

type Campsite struct {
	ID        uuid.UUID `json:"campsite_id" from:"campsite_id"`
	Name      string    `json:"name" from:"name"`
	Location  string    `json:"location" from:"location"`
	Latitude  float64   `json:"latitude" from:"latitude"`
	Lontitude float64   `json:"lontitude" from:"lontitude"`
	AreaCamp  string    `json:"area_camp" from:"area_camp"`
	Price     int       `json:"price" from:"price"`
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}
