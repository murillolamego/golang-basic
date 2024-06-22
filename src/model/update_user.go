package model

import "github.com/murillolamego/golang-basic/src/config/rest_err"

func (ud *UserDomain) UpdateUser(string) (*UserDomain, *rest_err.RestErr) {
	return ud, nil
}
