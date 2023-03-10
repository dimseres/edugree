package services

import (
	"authorization/internal/helpers"
	"authorization/internal/models"
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
	jwtPayload["user_id"] = user.Id
	jwtPayload["name"] = user.BaseUser.FullName

	if user.Role != nil {
		jwtPayload["role"] = user.Role.Id
		jwtPayload["role_name"] = user.Role.Slug
	}

	return helpers.CreateAuthToken(jwtPayload)
}
