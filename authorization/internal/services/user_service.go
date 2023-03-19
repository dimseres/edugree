package services

import "authorization/internal/models"

type UserRepository interface {
	GetUserById(id uint) (*models.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (self *UserService) GetUser(id uint) *models.User {
	user, _ := self.repository.GetUserById(id)
	return user
}

type RegistrationData struct {
	Email    string
	Password string
	Phone    string
	FullName string
	Avatar   *string
	Bio      *string
}
