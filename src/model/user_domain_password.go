package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptUserPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.password), 12)
	if err != nil {
		return err
	}

	ud.password = string(bytes)
	return nil
}
