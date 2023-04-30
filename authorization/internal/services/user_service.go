package services

import (
	"authorization/internal/dto"
	"authorization/internal/helpers"
	"authorization/internal/models"
)

type UserRepository interface {
	GetUserById(id uint) (*models.User, error)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
	CreateNewUser(user *models.User) (*models.User, error)
	GetUsersWithPagination(orgid uint, page int, perpage int) (*[]models.User, error)
}

type UserService struct {
	repository UserRepository
	BaseService
}

func NewUserService(repository UserRepository, tenantContext *TenantContext) UserService {
	return UserService{
		repository:  repository,
		BaseService: BaseService{tenantContext: tenantContext},
	}
}

func (self *UserService) GetUser(id uint) *models.User {
	user, _ := self.repository.GetUserById(id)
	return user
}

func (self *UserService) GetUsersWithPagination(page int, perpage int) (*[]models.User, error) {
	orgID := self.tenantContext.Id
	users, err := self.repository.GetUsersWithPagination(orgID, page, perpage)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (self *UserService) GetUserWith(id uint, relation *[]string) *models.User {
	var user models.User
	_, _ = self.repository.LoadRelation(&user, *relation...)
	return &user
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
