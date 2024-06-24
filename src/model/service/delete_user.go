package service

import "github.com/murillolamego/golang-basic/src/config/rest_err"

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	domainErr := ud.userRepository.DeleteUser(userId)
	if domainErr != nil {
		return domainErr
	}

	return nil
}
