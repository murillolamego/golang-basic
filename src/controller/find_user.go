package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")

		c.JSON(errorMessage.Code, errorMessage.Message)
		return
	}

	userDomain, err := uc.service.FindUserByID(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := rest_err.NewBadRequestError("Email is not valid")

		c.JSON(errorMessage.Code, errorMessage.Message)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
