package converter

import (
	"github.com/murillolamego/golang-basic/src/model"
	"github.com/murillolamego/golang-basic/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
