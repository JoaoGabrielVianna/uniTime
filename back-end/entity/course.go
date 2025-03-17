package entity

import "time"

type Course struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:text"`
	Years       []Year    `json:"years" gorm:"foreignKey:CourseID"`
	Users       []User    `json:"users" gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time `json:"created-at" gorm:"autoCreateTime"`
}
