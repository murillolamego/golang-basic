package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/validation"
	"github.com/murillolamego/golang-basic/src/controller/model/request"
	"github.com/murillolamego/golang-basic/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	user, err := domain.CreateUser()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, user)
}
