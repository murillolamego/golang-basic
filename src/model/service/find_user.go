package service

import (
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"github.com/murillolamego/golang-basic/src/model"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserByEmail(email)
}
