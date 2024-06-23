package main

import (
	"github.com/murillolamego/golang-basic/src/controller"
	"github.com/murillolamego/golang-basic/src/model/repository"
	"github.com/murillolamego/golang-basic/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
