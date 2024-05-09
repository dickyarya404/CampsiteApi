package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Facility struct {
	*gorm.Model
	ID          uuid.UUID `json:"id" from:"id" gorm:"primary_key"`
	CampsiteID  uuid.UUID `json:"campsite_id" from:"campsite_id" gorm:"primary_key"`
	Name        string    `json:"name" from:"name"`
	Description string    `json:"description" from:"description"`
	CreatedAt   time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time `gorm:"type:DATETIME(0)"`
}
