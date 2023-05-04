package repositories

import (
	"authorization/internal/database"
	"authorization/internal/database/helpers"
	"authorization/internal/models"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

type RepositoryInterface interface {
}

type UserRepository struct {
	BaseRepositoryHelpers
	cache CacheRepository
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
}

func NewUserRepository() UserRepository {
	return UserRepository{
		BaseRepositoryHelpers{
			db: database.GetConnection(),
		},
		NewCacheRepository(),
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
	res := rep.db.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (rep *UserRepository) GetUsersWithPagination(orgId uint, page int, perPage int) (*[]models.User, int64, error) {
	pagination := PaginationConfig{
		Page:    page,
		PerPage: perPage,
	}
	var membership []models.Membership
	var users []models.User
	var total int64

	//var result []struct {
	//	Id                uint64
	//	Email             string            `gorm:"not null;unique" json:"email"`
	//	Password          string            `gorm:"not null" json:"-"`
	//	PasswordResetCode *string           `gorm:"size:256" json:"-"`
	//	Phone             string            `gorm:"size:256;not null;unique" json:"phone"`
	//	FullName          string            `gorm:"size:512;not null" json:"full_name"`
	//	Avatar            *string           `json:"avatar"`
	//	Bio               *string           `gorm:"type:text" json:"bio"`
	//	Active            bool              `gorm:"not null;default:true" json:"active"`
	//	Membership        models.Membership `json:"membership"`
	//}

	// TODO сделать запрос из members и сопоставить роль и пользователя в ответе
	//result := map[string]interface{}{}

	res := rep.db.Scopes(rep.Paginate(&pagination)).Joins("User").Joins("Role").Find(&membership, "memberships.organization_id = ?", orgId)
	if res.Error != nil {
		return nil, 0, res.Error
	}

	for _, member := range membership {
		member.User.DomainRole = member.Role
		users = append(users, member.User)
	}

	return &users, total, nil
}

func (rep *UserRepository) CreateNewUser(payload *models.User) (*models.User, error) {
	var user models.User
	_ = rep.db.Where("email = ? or phone = ?", payload.Email, payload.Phone).First(&user)
	if user.Id > 0 {
		return nil, errors.New("user already registered")
	}
	data := rep.db.Create(&payload)
	if data.Error != nil {
		return nil, rep.handleError(data.Error)
	}
	return payload, nil
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
