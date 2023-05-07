package repositories

import (
	"authorization/internal/database"
	"authorization/internal/models"
)

type RoleRepository struct {
	BaseRepositoryHelpers
	cache CacheRepository
}

func NewRoleRepository() RoleRepository {
	return RoleRepository{
		BaseRepositoryHelpers{db: database.GetConnection()},
		NewCacheRepository(),
	}
}

func (self *RoleRepository) GetRolesBySlug(needle []string) (*[]models.Role, error) {
	var roles []models.Role
	res := self.db.Where("slug in ?", needle).Find(&roles)
	if res.Error != nil {
		return nil, res.Error
	}
	return &roles, nil
}
