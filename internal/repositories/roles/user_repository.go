package users

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/database/helpers"
	"edugree_auth/internal/database/models"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func NewRepository() RoleRepository {
	return RoleRepository{
		db: database.GetConnection(),
	}
}

func (rep *RoleRepository) payloadToModel(payload *RoleDataPayload) models.Role {
	return models.Role{
		Title:  payload.Title,
		Domain: payload.Domain,
		Slug:   payload.Slug,
	}
}

func (rep *RoleRepository) handleError(err error) error {
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

func (rep *RoleRepository) GetUserById(id uint) *models.Role {
	user := models.Role{}
	rep.db.Preload("Role").First(&user, id)
	return &user
}

func (rep *RoleRepository) CreateNewUser(payload *RoleDataPayload) (error, *models.Role) {
	newUser := rep.payloadToModel(payload)
	data := rep.db.Create(&newUser)
	if data.Error != nil {
		return rep.handleError(data.Error), nil
	}
	return nil, &newUser
}

func (rep *RoleRepository) CreateNewUsers(payload []*RoleDataPayload) (bool, []models.Role) {
	var inserted []models.Role
	for i := 0; i < len(payload); i++ {
		inserted = append(inserted, rep.payloadToModel(payload[i]))
	}
	data := rep.db.Create(&inserted)
	if data == nil {
		return false, inserted
	}
	return true, inserted
}

func (rep *RoleRepository) UpdateUser(id uint, payload *RoleDataPayload) (error, *models.Role) {
	user := models.Role{}
	user.Id = id
	updatedUser := rep.payloadToModel(payload)
	queryResult := rep.db.Model(&user).Updates(&updatedUser)
	if queryResult.Error != nil {
		return rep.handleError(queryResult.Error), nil
	}

	return nil, &user
}

func (rep *RoleRepository) DeleteUser(id uint) (error, bool) {
	user := models.Role{}
	result := rep.db.Delete(&user, id)
	if result.Error != nil {
		return rep.handleError(result.Error), false
	}
	return nil, true
}
