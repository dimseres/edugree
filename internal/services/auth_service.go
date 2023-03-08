package services

import (
	"edugree_auth/internal/helpers"
	"edugree_auth/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt"
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

func (self *AuthService) SignIn(email string, password string) (error, *models.User) {
	err, userData := self.repository.CheckLoginData(email)
	if err != nil {
		return err, nil
	}

	err = helpers.ComparePasswordAndHash(password, userData.Password)
	if err != nil {
		return err, nil
	}

	return nil, userData
}

func (self *AuthService) CreateJwtToken(user *models.User) (error, string) {
	jwtPayload := make(map[string]interface{})
	//jwtPayload := map[string]interface{}{
	//	"user_id": user.BaseUser.Id,
	//	"name":    user.BaseUser.FullName,
	//	"role":    user.Role.Slug,
	//	"role_id": user.RoleId,
	//}
	fmt.Println()
	jwtPayload["user_id"] = user.BaseUser.Id
	jwtPayload["name"] = user.BaseUser.FullName
	return helpers.CreateAuthToken(jwtPayload)
}
