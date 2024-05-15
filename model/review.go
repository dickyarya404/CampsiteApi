package model

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID `json:"id" from:"id" gorm:"primary_key"`
	Rating      int       `json:"rating" from:"rating"`
	Comment     string    `json:"comment" from:"comment"`
	Review_date string    `json:"review_date" from:"review_date"`
	CreatedAt   time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time `gorm:"type:DATETIME(0)"`
}
