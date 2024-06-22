package view

import (
	"github.com/murillolamego/golang-basic/src/controller/model/response"
	"github.com/murillolamego/golang-basic/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "test",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
