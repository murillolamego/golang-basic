package service

import (
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	domainErr := ud.userRepository.UpdateUser(userId, userDomain)
	if domainErr != nil {
		return domainErr
	}

	return nil
}
