package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Student Role = "student"
	Teacher Role = "teacher"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Email     string    `json:"email" gorm:"unique;type:varchar(100)"`
	Password  string    `json:"-" gorm:"type:varchar(100)"`
	Role      Role      `json:"role" gorm:"type:varchar(100)"`
	CourseID  uuid.UUID `json:"course_id" gorm:"type:uuid;index"` // Relacionado diretamente à tabela courses	Year      int       `json:"year" gorm:"type:int"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
