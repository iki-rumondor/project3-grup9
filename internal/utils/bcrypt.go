package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	salt := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(h, p string) error {
	hash, pass := []byte(h), []byte(p)

	if err := bcrypt.CompareHashAndPassword(hash, pass); err != nil {
		return err
	}

	return nil
}
