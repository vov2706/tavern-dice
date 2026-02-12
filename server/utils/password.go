package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(p string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), 8)

	return string(hash)
}

func IsValidPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
