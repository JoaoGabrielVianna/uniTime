package entity

import "time"

type CalendarEvent struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	CalendarID  string    `json:"calendar_id" gorm:"index;foreignKey:CalendarID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string    `json:"title" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	EventDate   time.Time `json:"event_date" gorm:"type:date"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	AuthorID    string    `json:"author_id" gorm:"index;foreignKey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
