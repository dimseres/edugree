package services

import (
	"edugree_auth/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	CheckLoginData(email string) (error, *models.User)
}

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"user_name"`
}

type AuthService struct {
	repository AuthRepository
}

func NewAuthService(repository AuthRepository) AuthService {
	return AuthService{
		repository: repository,
	}
}

func (self *AuthService) SignIn(email string, password string) (error, interface{}) {
	err, userData := self.repository.CheckLoginData(email)

	fmt.Println(err, password)

	if err != nil {
		return err, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userData.Password))

	if err != nil {
		return err, nil
	}

	return nil, userData
}
