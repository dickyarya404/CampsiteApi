package model

import (
	"time"

	"github.com/google/uuid"
)

type Availability struct {
	ID        uuid.UUID ` gorm:"primary_key" json:"availability_id" from:"availability_id"`
	Quantity  int       `json:"quantity" from:"quantity"`
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}
