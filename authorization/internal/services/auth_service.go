package services

import (
	"authorization/internal/constants"
	"authorization/internal/helpers"
	"authorization/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

// REFRESH_LIFETIME One week
type AuthRepository interface {
	CheckLoginData(email string) (error, *models.User)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
	RegisterRefreshToken(user *models.User, token string, salt string, lifeTime time.Duration) error
	GetRefreshToken(token string) (*models.Token, error)
	DeleteRefreshToken(token *models.Token) error
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

	_, err = self.repository.LoadRelation(userData, "Membership",
		"Membership.Role",
		"Membership.Organization.Services")

	if err != nil {
		return err, nil
	}

	return nil, userData
}

func (self *AuthService) CreateJwtToken(user *models.User, domain string) (error, string) {
	jwtPayload := helpers.JwtData{}
	fmt.Println()

	var memberships []helpers.JwtMembershipData

	for _, member := range user.Membership {
		var allowedServices []string

		if member.Organization.Services != nil {
			for _, member := range *member.Organization.Services {
				allowedServices = append(allowedServices, member.Slug)
			}
		}

		var membershipRole helpers.JwtMembershipData

		if member.Role != nil {
			membershipRole.Role = &member.Role.Slug
		}

		if member.Organization != nil {
			membershipRole.Organization = &member.Organization.Domain
			membershipRole.OrganizationId = &member.OrganizationId
		}

		if member.Organization != nil {
			membershipRole.Organization = &member.Organization.Domain
		}

		if len(allowedServices) > 0 {
			membershipRole.Services = &allowedServices
		}

		memberships = append(memberships, membershipRole)
	}

	jwtPayload.UserId = user.Id
	jwtPayload.Name = user.FullName
	jwtPayload.Membership = memberships

	return helpers.CreateAuthToken(jwtPayload)
}

func (self *AuthService) CreateRefreshToken(token string, user *models.User) (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	token, err = helpers.CreatePasswordHash(token + uid.String())
	err = self.repository.RegisterRefreshToken(user, token, uid.String(), constants.REFRESH_LIFETIME)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (self *AuthService) Logout(token string, refresh string) error {
	dbToken, err := self.repository.GetRefreshToken(refresh)
	if err != nil {
		return err
	}

	err = helpers.ComparePasswordAndHash(token+dbToken.Salt, dbToken.Token)
	if err != nil {
		return err
	}

	err = self.repository.DeleteRefreshToken(dbToken)
	if err != nil {
		return err
	}

	return nil
}

type GenerateTokenDTO struct {
	User    *models.User
	Token   string
	Refresh string
}

func (self *AuthService) GenerateTokenFromRefresh(refreshToken string) (*GenerateTokenDTO, error) {
	res, err := self.repository.GetRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	_, err = self.repository.LoadRelation(res, "User", "User.Membership", "User.Membership.Organization", "User.Membership.Role")
	if err != nil {
		return nil, err
	}

	err, token := self.CreateJwtToken(&res.User, "")
	if err != nil {
		return nil, err
	}

	newRefresh, err := self.CreateRefreshToken(token, &res.User)
	if err != nil {
		return nil, err
	}

	user := res.User

	err = self.repository.DeleteRefreshToken(res)
	if err != nil {
		return nil, err
	}

	return &GenerateTokenDTO{
		User:    &user,
		Token:   token,
		Refresh: newRefresh,
	}, nil

}
