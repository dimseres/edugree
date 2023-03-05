package repositories

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/models"
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

func (self *AuthRepository) CheckLoginData(email string, password string) (error, *models.User) {
	user := models.User{}
	res := self.db.First(&user, "email = ? and password = ?", email, password)
	if res.Error != nil {
		return res.Error, nil
	}
	return nil, &user
}
