package service

import (
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	err := userDomain.EncryptUserPassword()
	if err != nil {

		logger.Error("error trying to encrypt user password", err, zap.String("journey", "createUser"))
		restErr := rest_err.NewInternalServerErrorfunc("error trying to encrypt user password")

		return userDomain, restErr
	}

	return userDomain, nil
}
