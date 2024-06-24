package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/validation"
	"github.com/murillolamego/golang-basic/src/controller/model/request"
	"github.com/murillolamego/golang-basic/src/model"
	"github.com/murillolamego/golang-basic/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("init LoginUser controller", zap.String("journey", "loginUser"))

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err, zap.String("journey", "loginUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	user, token, err := uc.service.LoginUser(domain)
	if err != nil {
		logger.Error("error trying to call LoginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user logged in successfully",
		zap.String("userId", domain.GetID()), zap.String("journey", "loginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(user))
}
