package models

import "time"

type Location struct {
	ID          uint       `gorm:"primary_key" json:"id"`
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
