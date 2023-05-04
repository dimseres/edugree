package repositories

import (
	"authorization/internal/constants"
	"authorization/internal/database"
	"authorization/internal/models"
	"errors"
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

func (self *MembershipRepository) DeleteMember(memberId uint, organizationId uint) error {
	var member models.Membership
	resp := self.db.Joins("Role").Where("user_id = ?", memberId).Take(&member)
	if resp.Error != nil {
		return resp.Error
	}

	if member.Role != nil {
		if member.Role.Slug == constants.SubOwner {
			return errors.New("cant delete owner")
		}
	}

	resp = self.db.Delete(&member, "user_id = ?", memberId)
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
