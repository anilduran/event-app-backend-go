package models

import "time"

type User struct {
	ID                 uint   `gorm:"primary_key" json:"id"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	Password           string
	ProfilePhotoUrl    string   `json:"profile_photo_url"`
	Events             []Event  `gorm:"foreignkey:CreatorID"`
	ParticipatedEvents []*Event `gorm:"many2many:event_participants;"`
	Comments           []Comment
	Locations          []Location `gorm:"foreignkey:CreatorID"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
}
