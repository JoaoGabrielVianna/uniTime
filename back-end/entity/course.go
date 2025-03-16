package entity

import "time"

type Year struct {
	ID       string `json:"id"`
	Year     int    `json:"year"`
	Teachers []User `json:"teachers"`
	Students []User `json:"students"`
}
type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Years       []Year    `json:"years"`
	CreatedAt   time.Time `json:"created-at"`
}
