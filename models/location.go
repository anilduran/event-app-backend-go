package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"image_url"`
	CreatorID   uint       `json:"creator_id"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Events      []Event    `gorm:"foreignkey:LocationID"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (location *Location) BeforeCreate(tx *gorm.DB) (err error) {
	location.ID = uuid.New()
	return
}
