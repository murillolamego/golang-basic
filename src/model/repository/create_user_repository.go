package repository

import (
	"context"
	"os"

	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
	"github.com/murillolamego/golang-basic/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init createUser repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerErrorfunc(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
