package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hash, password string) error {
	if hash == "" {
		return fmt.Errorf("hash parameter is empty")
	}
	if password == "" {
		return fmt.Errorf("password is empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
