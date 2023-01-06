package users

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/database/helpers"
	"edugree_auth/internal/database/models"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"reflect"
)

func NewRepository() Repository {
	return Repository{
		db: database.GetConnection(),
	}
}

func (rep *Repository) payloadToModel(payload *UserDataPayload) models.User {
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

func handleError(err error) error {
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

func ForeachStruct(data interface{}, handler func(key any, value any, idx int)) {
	if data == nil {
		return
	}
	values := reflect.ValueOf(data)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		handler(types.Field(i).Name, values.Field(i), i)
	}
}

func (rep *Repository) GetUserById(id uint) *models.User {
	user := models.User{}
	rep.db.Preload("Role").First(&user, id)
	return &user
}

func (rep *Repository) CreateNewUser(payload *UserDataPayload) (error, *models.User) {
	newUser := rep.payloadToModel(payload)
	data := rep.db.Create(&newUser)
	if data.Error != nil {
		return handleError(data.Error), nil
	}
	return nil, &newUser
}

func (rep *Repository) CreateNewUsers(payload []*UserDataPayload) (bool, []models.User) {
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

func (rep *Repository) UpdateUser(id uint, payload *UserDataPayload) (error, *models.User) {
	user := models.User{}
	user.Id = id
	updatedUser := rep.payloadToModel(payload)
	queryResult := rep.db.Model(&user).Updates(&updatedUser)
	if queryResult.Error != nil {
		return handleError(queryResult.Error), nil
	}

	return nil, nil
}

func (rep *Repository) DeleteUser(id uint) (error, bool) {
	user := models.User{}
	result := rep.db.Delete(&user, id)
	if result.Error != nil {
		return handleError(result.Error), false
	}
	return nil, true
}
