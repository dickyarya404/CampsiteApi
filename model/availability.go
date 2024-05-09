package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Availability struct {
	*gorm.Model
	ID         uuid.UUID ` gorm:"primary_key" json:"availability_id" from:"availability_id"`
	CampsiteID uuid.UUID `gorm:"foreign_key" json:"campsite_id" from:"campsite_id" `
	Quantity   int       `json:"quantity" from:"quantity"`
	CreatedAt  time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt  time.Time `gorm:"type:DATETIME(0)"`
}
