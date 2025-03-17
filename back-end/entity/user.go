package entity

import "time"

type Role string

const (
	Student Role = "student"
	Teacher Role = "teacher"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Email     string    `json:"email" gorm:"unique;type:varchar(100)"`
	Password  string    `json:"-" gorm:"type:varchar(100)"`
	Role      Role      `json:"role" gorm:"type:varchar(100)"`
	Course    *Course   `gorm:"foreignKey:CourseID" json:"course"`
	CourseID  string    `json:"course_id" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
