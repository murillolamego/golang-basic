package service

import (
	"github.com/murillolamego/golang-basic/src/config/logger"
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	user, restErr := ud.FindUserByEmail(userDomain.GetEmail())
	if restErr != nil {
		logger.Error("invalid credentials provided, please try again", restErr, zap.String("journey", "loginUser"))

		restErr := rest_err.NewBadRequestError("invalid credentials provided, please try again")

		return nil, restErr
	}

	err := userDomain.CheckUserPassword(user.GetPassword())
	if err != nil {
		logger.Error("invalid credentials provided, please try again", err, zap.String("journey", "loginUser"))

		restErr := rest_err.NewBadRequestError("invalid credentials provided, please try again")

		return nil, restErr
	}

	userDomainRepository, domainErr := ud.userRepository.CreateUser(userDomain)
	if domainErr != nil {
		return nil, domainErr
	}

	return userDomainRepository, nil
}
