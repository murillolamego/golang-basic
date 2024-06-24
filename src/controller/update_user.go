package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/config/validation"
	"github.com/murillolamego/golang-basic/src/controller/model/request"
	"github.com/murillolamego/golang-basic/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("init UpdateUser controller", zap.String("journey", "updateUser"))

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err, zap.String("journey", "updateUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a valid hex value")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUpdateUserDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("error trying to call UpdateUser service", err, zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user updated successfully",
		zap.String("userId", userId), zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
