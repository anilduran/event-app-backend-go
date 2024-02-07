package main

import (
	"context"

	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/routes"
	"example.com/event-app-backend-go/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("unable to load SDK config")
	}

	client := s3.NewFromConfig(cfg)

	utils.InitializePresigner(client)

	err = db.InitializeDB()

	if err != nil {
		panic("Could not connect to the database")
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")

}
