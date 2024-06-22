package model

import (
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() (*UserDomain, *rest_err.RestErr) {

	err := ud.EncryptUserPassword()
	if err != nil {

		logger.Error("error trying to encrypt user password", err, zap.String("journey", "createUser"))
		restErr := rest_err.NewInternalServerErrorfunc("error trying to encrypt user password")

		return ud, restErr
	}

	return ud, nil
}
