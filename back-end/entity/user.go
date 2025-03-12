package entity

import "time"

type Role string

const (
	Student Role = "student"
	Teacher Role = "teacher"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Course    string    `json:"course"`
}
