package models

import "time"

type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Content   string     `json:"content"`
	UserID    uint       `json:"user_id"`
	EventID   uint       `json:"event_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
