package repositories

import (
	"authorization/internal/constants"
	"authorization/internal/database"
	"authorization/internal/models"
	"authorization/internal/transport/rest/forms"
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

func (self *MembershipRepository) InviteMembers(members []forms.MemberInviteForm, organizationId uint) error {
	lim := 100
	var chunk []forms.MemberInviteForm
	for idx, member := range members {
		chunk = append(chunk, member)
		if idx%lim == 0 {
			self.inviteMemberFromMail(chunk)
			chunk = make([]forms.MemberInviteForm, 1)
		}
	}

	return nil
}

func (self *MembershipRepository) inviteMemberFromMail(members []forms.MemberInviteForm) error {
	var emails, roles []string
	for _, member := range members {
		emails = append(emails, member.Email)
		roles = append(roles, member.Role)
	}
	var users []models.User
	res := self.db.Where("email IN ?", emails).Find(&users)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
