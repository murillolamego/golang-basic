package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/murillolamego/golang-basic/src/config/database/mongodb"
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/controller/routes"
)

func main() {
	logger.Info("About to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database, error=%s", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
