package services

import (
	"edugree_auth/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt"
)

type AuthRepository interface {
	CheckLoginData(email string, password string) (error, *models.User)
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
	err, userData := self.repository.CheckLoginData(email, password)

	fmt.Println(err)

	if err != nil {
		return err, nil
	}

	return nil, userData
}
