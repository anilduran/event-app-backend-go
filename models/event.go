package models

import "time"

type Event struct {
	ID           uint        `gorm:"primary_key" json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Capacity     uint        `json:"capacity"`
	CreatorID    uint        `json:"creator_id"`
	Categories   []*Category `gorm:"many2many:event_categories;"`
	Participants []*User     `gorm:"many2many:event_participants;"`
	CreatedAt    *time.Time  `json:"created_at"`
	UpdatedAt    *time.Time  `json:"updated_at"`
	DeletedAt    *time.Time  `json:"deleted_at"`
}
