package model

import (
	"time"

	"github.com/google/uuid"
)

type Facility struct {
	ID          uuid.UUID `json:"id" from:"id" gorm:"primary_key"`
	Name        string    `json:"name" from:"name"`
	Description string    `json:"description" from:"description"`
	CreatedAt   time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time `gorm:"type:DATETIME(0)"`
}
