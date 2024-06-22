package model

import (
	"github.com/murillolamego/golang-basic/src/config/rest_err"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptUserPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.Password), 12)
	if err != nil {
		return err
	}

	ud.Password = string(bytes)
	return nil
}

type UserDomainInterface interface {
	CreateUser() (*UserDomain, *rest_err.RestErr)
	FindUser(string) *rest_err.RestErr
	UpdateUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
