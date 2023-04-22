package repositories

import (
	"authorization/internal/database"
	"authorization/internal/models"
)

type MembershipRepository struct {
	BaseRepositoryHelpers
}

func NewMembershipRepository() MembershipRepository {
	return MembershipRepository{
		BaseRepositoryHelpers{
			db: database.GetConnection(),
		},
	}
}

func (self *MembershipRepository) GetMembershipData(organizationId uint, userId uint) (*models.Membership, error) {
	var member models.Membership
	res := self.db.Where("user_id = ? and organization_id = ?", userId, organizationId).
		Preload("Role").
		Preload("Organization").
		Preload("Organization.Services").
		First(&member)
	if res.Error != nil {
		return nil, res.Error
	}

	return &member, nil
}
