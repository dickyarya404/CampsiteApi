package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	*gorm.Model
	ID        uuid.UUID `gorm:"primary_key" json:"admin_id" from:"admin_id"`
	Name      string    `json:"name" from:"name"`
	Password  string    `json:"password" from:"password"`
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}
