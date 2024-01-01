package models

import "time"

type Category struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Events    []*Event   `gorm:"many2many:event_categories;"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
