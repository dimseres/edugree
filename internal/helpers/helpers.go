package helpers

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"os"
)

func CreatePasswordHash(password string) (string, error) {
	firstSalt, err := scrypt.Key([]byte(password), []byte(os.Getenv("SALT")), 2048, 4, 2, 32)
	passwordHash, err := scrypt.Key([]byte(password), firstSalt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(passwordHash), nil
}

type CompareFailed struct {
	Message string
}

func (e *CompareFailed) Error() string {
	return fmt.Sprintf("parse %v:", e.Message)
}

func ComparePasswordAndHash(password string, hashedPassword string) error {
	password, err := CreatePasswordHash(password)
	if err != nil {
		return err
	}
	if password != hashedPassword {
		return &CompareFailed{Message: "password not compared"}
	}
	return nil
}
