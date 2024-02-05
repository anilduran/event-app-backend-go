package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID           uuid.UUID   `gorm:"type:uuid;primary_key;"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	ImageUrl     string      `json:"image_url"`
	Capacity     uint        `json:"capacity"`
	CreatorID    uuid.UUID   `json:"creator_id"`
	Categories   []*Category `gorm:"many2many:event_categories;"`
	Participants []*User     `gorm:"many2many:event_participants;"`
	Comments     []Comment
	LocationID   uuid.UUID  `json:"location_id"`
	Events       []Event    `gorm:"foreignkey:LocationID"`
	EventDate    *time.Time `json:"event_date"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func (event *Event) BeforeCreate(tx *gorm.DB) (err error) {
	event.ID = uuid.New()
	return
}
