package model

import (
	"time"
)

type Admin struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
}
