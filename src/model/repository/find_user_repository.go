package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
	"github.com/murillolamego/golang-basic/src/model/repository/entity"
	"github.com/murillolamego/golang-basic/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUserByID repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex((id))
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with ID %s not found", id)

			return nil, rest_err.NewNotFoundErrorfunc(errorMessage)
		}

		return nil, rest_err.NewInternalServerErrorfunc("Error trying to find user by ID")
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUserByEmail repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", email)

			return nil, rest_err.NewNotFoundErrorfunc(errorMessage)
		}

		return nil, rest_err.NewInternalServerErrorfunc("Error trying to find user by email")
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}
