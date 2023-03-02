package users

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/database/helpers"
	"edugree_auth/internal/database/models"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func NewRepository() UserRepository {
	return UserRepository{
		db: database.GetConnection(),
	}
}

func (rep *UserRepository) payloadToModel(payload *UserDataPayload) models.User {
	return models.User{
		Email:             payload.Email,
		Password:          payload.Password,
		PasswordResetCode: payload.PasswordResetCode,
		Phone:             payload.Phone,
		FullName:          payload.FullName,
		Avatar:            payload.Avatar,
		Bio:               payload.Bio,
		Active:            payload.Active,
		RoleId:            payload.RoleId,
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

func (rep *UserRepository) GetUserById(id uint) *models.User {
	user := models.User{}
	rep.db.Preload("Role").First(&user, id)
	return &user
}

func (rep *UserRepository) CreateNewUser(payload *UserDataPayload) (error, *models.User) {
	newUser := rep.payloadToModel(payload)
	data := rep.db.Create(&newUser)
	if data.Error != nil {
		return rep.handleError(data.Error), nil
	}
	return nil, &newUser
}

func (rep *UserRepository) CreateNewUsers(payload []*UserDataPayload) (bool, []models.User) {
	var inserted []models.User
	for i := 0; i < len(payload); i++ {
		inserted = append(inserted, rep.payloadToModel(payload[i]))
	}
	data := rep.db.Create(&inserted)
	if data == nil {
		return false, inserted
	}
	return true, inserted
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
