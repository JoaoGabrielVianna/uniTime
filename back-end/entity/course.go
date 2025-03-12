package entity

import "time"

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Teachers    []User    `json:"teachers"`
	Students    []User    `json:"students"`
	CreatedAt   time.Time `json:"created-at"`
}
