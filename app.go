package main

import (
	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	err := db.InitializeDB()

	if err != nil {
		panic("Could not connect to the database")
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")

}
