package users

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/database/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return Repository{
		db: database.GetConnection(),
	}
}

func (rep *Repository) GetUserById(id int) *models.User {
	user := models.User{}
	rep.db.Preload("Role").First(&user, id)
	return &user
}
