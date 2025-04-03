package entity

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Period    string    `json:"period" gorm:"type:varchar(100)"`
	Year      int       `json:"year" gorm:"type:int"`
	Acronym   string    `json:"acronym" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
