package entity

import "github.com/google/uuid"

type Course_year struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	CourseID string    `json:"course_id" gorm:"index;foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Year     int       `json:"year" gorm:"type:int"`
}
