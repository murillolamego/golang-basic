package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/controller"
	"github.com/murillolamego/golang-basic/src/controller/routes"
	"github.com/murillolamego/golang-basic/src/model/service"
)

func main() {
	logger.Info("About to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
