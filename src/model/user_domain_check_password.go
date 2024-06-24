package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) CheckUserPassword(hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(ud.password))
	if err != nil {
		return err
	}

	return nil
}
