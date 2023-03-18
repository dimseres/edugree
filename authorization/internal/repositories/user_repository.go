package repositories

import (
	"authorization/internal/database"
	"authorization/internal/database/helpers"
	"authorization/internal/models"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
}

type UserRepository struct {
	db *gorm.DB
}

type UserDataPayload struct {
	Email             string
	Password          string
	PasswordResetCode *string
	Phone             string
	FullName          string
	Avatar            *string
	Bio               *string
	Active            bool
	RoleId            *uint
}

func NewUserRepository() UserRepository {
	return UserRepository{
		db: database.GetConnection(),
	}
}

func (rep *UserRepository) payloadToModel(payload *UserDataPayload) models.User {
	return models.User{
		Email:             payload.Email,
		Phone:             payload.Phone,
		FullName:          payload.FullName,
		Avatar:            payload.Avatar,
		Bio:               payload.Bio,
		Active:            payload.Active,
		Password:          payload.Password,
		PasswordResetCode: payload.PasswordResetCode,
	}
}

func (rep *UserRepository) handleError(err error) error {
	if err != nil {
		switch err.(type) {
		default:
			return errors.New(err.Error())
		case *pgconn.PgError:
			return errors.New(helpers.GetErrorMessage(err.(*pgconn.PgError)))
		}
	}
	return nil
}

func (rep *UserRepository) GetUserById(id uint) (*models.User, error) {
	user := models.User{}
	res := rep.db.Preload("Membership").Preload("DomainRole").First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (rep *UserRepository) CreateNewUser(payload *UserDataPayload) (error, *models.User) {
	newUser := rep.payloadToModel(payload)
	data := rep.db.Create(&newUser)
	if data.Error != nil {
		return rep.handleError(data.Error), nil
	}
	return nil, &newUser
}

func (rep *UserRepository) CreateNewUsers(payload []*UserDataPayload) (bool, *[]models.User) {
	var inserted []models.User
	for i := 0; i < len(payload); i++ {
		inserted = append(inserted, rep.payloadToModel(payload[i]))
	}
	data := rep.db.Create(&inserted)
	if data == nil {
		return false, &inserted
	}
	return true, &inserted
}

func (rep *UserRepository) UpdateUser(id uint, payload *UserDataPayload) (error, *models.User) {
	user := models.User{}
	user.Id = id
	updatedUser := rep.payloadToModel(payload)
	queryResult := rep.db.Model(&user).Updates(&updatedUser)
	if queryResult.Error != nil {
		return rep.handleError(queryResult.Error), nil
	}

	return nil, &user
}

func (rep *UserRepository) DeleteUser(id uint) (error, bool) {
	user := models.User{}
	result := rep.db.Delete(&user, id)
	if result.Error != nil {
		return rep.handleError(result.Error), false
	}
	return nil, true
}
