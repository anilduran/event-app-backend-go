package db

import (
	"os"

	"example.com/event-app-backend-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() (err error) {

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}

	err = DB.AutoMigrate(models.User{}, models.Event{}, models.Category{})

	return

}
