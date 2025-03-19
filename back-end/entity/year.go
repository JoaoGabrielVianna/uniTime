package entity

type Year struct {
	ID       string `json:"id" gorm:"primaryKey"`
	CourseID string `json:"course_id" gorm:"index"` // Chave estrangeira para Course
	Year     int    `json:"year" gorm:"type:int"`
	Teachers []User `json:"teachers" gorm:"many2many:year_teachers"`
	Students []User `json:"students" gorm:"many2many:year_students"`
}
