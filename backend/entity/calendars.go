package entity

import "time"

type Calendar struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	CourseYearID string    `json:"course_year_id" gorm:"index;foreignKey:CourseYearID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name         string    `json:"name" gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
