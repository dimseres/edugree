package services

import "edugree_auth/internal/models"

type UserRepository interface {
	GetUserById(id uint) *models.User
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (self *UserService) GetUser(id uint) *models.PublicUser {
	user := self.repository.GetUserById(id)
	return &models.PublicUser{
		BaseModel: user.BaseModel,
		BaseUser:  user.BaseUser,
		Role:      user.Role,
		Token:     user.Token,
	}
}
