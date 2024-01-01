package models

import "time"

type User struct {
	ID                 uint   `gorm:"primary_key" json:"id"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	Password           string
	Events             []Event    `gorm:"foreignkey:CreatorID"`
	ParticipatedEvents []*Event   `gorm:"many2many:event_participants;"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
}
