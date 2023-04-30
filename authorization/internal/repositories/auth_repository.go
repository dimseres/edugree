package repositories

import (
	"authorization/internal/database"
	"authorization/internal/models"
	"time"
)

type AuthRepository struct {
	BaseRepositoryHelpers
	cache CacheRepository
}

func NewAuthRepository() AuthRepository {
	return AuthRepository{
		BaseRepositoryHelpers{db: database.GetConnection()},
		NewCacheRepository(),
	}
}

func (self *AuthRepository) CheckLoginData(email string) (error, *models.User) {
	user := models.User{}
	res := self.db.First(&user, "email = ?", email)
	if res.Error != nil {
		return res.Error, nil
	}
	return nil, &user
}

func (self *AuthRepository) LoadUserRoleByDomain(user *models.User, domain string) error {
	organization := models.Organization{}
	res := self.db.Where(user).Find(&organization)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (self *AuthRepository) RegisterRefreshToken(user *models.User, token string, salt string, lifetime time.Duration) error {
	refresh := models.Token{
		UserId:   user.Id,
		Token:    token,
		Salt:     salt,
		ExpireAt: time.Now().Add(lifetime),
		User:     models.User{},
	}
	res := self.db.Create(&refresh)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (self *AuthRepository) GetRefreshToken(token string) (*models.Token, error) {
	refresh := models.Token{}
	res := self.db.Where("token = ? and expire_at > ?", token, time.Now().Format(time.DateTime)).First(&refresh)
	if res.Error != nil {
		return nil, res.Error
	}
	return &refresh, nil
}

func (self *AuthRepository) DeleteRefreshToken(token *models.Token) error {
	res := self.db.Delete(token)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (self *AuthRepository) DeleteRefreshTokenByUserId(userId uint) error {
	refresh := models.Token{}
	res := self.db.Where("user_id = ", userId).Delete(&refresh)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
