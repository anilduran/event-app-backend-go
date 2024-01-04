package models

import "time"

type Event struct {
	ID           uint        `gorm:"primary_key" json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	ImageUrl     string      `json:"image_url"`
	Capacity     uint        `json:"capacity"`
	CreatorID    uint        `json:"creator_id"`
	Categories   []*Category `gorm:"many2many:event_categories;"`
	Participants []*User     `gorm:"many2many:event_participants;"`
	Comments     []Comment
	LocationID   uint       `json:"location_id"`
	Events       []Event    `gorm:"foreignkey:LocationID"`
	EventDate    *time.Time `json:"event_date"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
