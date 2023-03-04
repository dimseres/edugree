package repositories

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/database/helpers"
	"edugree_auth/internal/models"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
}

type Repository struct {
	db *gorm.DB
}

type DataPayload struct {
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

func NewRepository() Repository {
	return Repository{
		db: database.GetConnection(),
	}
}

func (rep *Repository) payloadToModel(payload *DataPayload) models.User {
	return models.User{
		BaseUser: models.BaseUser{
			Email:    payload.Email,
			Phone:    payload.Phone,
			FullName: payload.FullName,
			Avatar:   payload.Avatar,
			Bio:      payload.Bio,
			Active:   payload.Active,
			RoleId:   payload.RoleId,
		},
		Password:          payload.Password,
		PasswordResetCode: payload.PasswordResetCode,
	}
}

func (rep *Repository) handleError(err error) error {
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

func (rep *Repository) GetUserById(id uint) *models.User {
	user := models.User{}
	rep.db.Preload("Role").Preload("Token").First(&user, id)
	return &user
}

func (rep *Repository) CreateNewUser(payload *DataPayload) (error, *models.User) {
	newUser := rep.payloadToModel(payload)
	data := rep.db.Create(&newUser)
	if data.Error != nil {
		return rep.handleError(data.Error), nil
	}
	return nil, &newUser
}

func (rep *Repository) CreateNewUsers(payload []*DataPayload) (bool, *[]models.User) {
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

func (rep *Repository) UpdateUser(id uint, payload *DataPayload) (error, *models.User) {
	user := models.User{}
	user.Id = id
	updatedUser := rep.payloadToModel(payload)
	queryResult := rep.db.Model(&user).Updates(&updatedUser)
	if queryResult.Error != nil {
		return rep.handleError(queryResult.Error), nil
	}

	return nil, &user
}

func (rep *Repository) DeleteUser(id uint) (error, bool) {
	user := models.User{}
	result := rep.db.Delete(&user, id)
	if result.Error != nil {
		return rep.handleError(result.Error), false
	}
	return nil, true
}
