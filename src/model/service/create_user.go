package service

import (
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	user, _ := ud.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already in use")
	}

	err := userDomain.EncryptUserPassword()
	if err != nil {
		logger.Error("error trying to encrypt user password", err, zap.String("journey", "createUser"))

		restErr := rest_err.NewInternalServerErrorfunc("error trying to encrypt user password")

		return userDomain, restErr
	}

	userDomainRepository, domainErr := ud.userRepository.CreateUser(userDomain)
	if domainErr != nil {
		return nil, domainErr
	}

	return userDomainRepository, nil
}
