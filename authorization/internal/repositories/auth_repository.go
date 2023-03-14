package repositories

import (
	"authorization/internal/database"
	"authorization/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	return AuthRepository{
		db: database.GetConnection(),
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
