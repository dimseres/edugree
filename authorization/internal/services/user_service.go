package services

import (
	"authorization/internal/dto"
	"authorization/internal/helpers"
	"authorization/internal/models"
)

type UserRepository interface {
	GetUserById(id uint) (*models.User, error)
	CreateNewUser(user *models.User) (*models.User, error)
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

func (self *UserService) CreateUser(userDTO *dto.CreateUserDTO) (*models.User, error) {
	passwordHash, err := helpers.CreatePasswordHash(userDTO.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    userDTO.Email,
		Password: passwordHash,
		Phone:    userDTO.Phone,
		FullName: userDTO.FullName,
		Active:   true,
	}

	savedUser, err := self.repository.CreateNewUser(&user)
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

type RegistrationData struct {
	Email    string
	Password string
	Phone    string
	FullName string
	Avatar   *string
	Bio      *string
}
