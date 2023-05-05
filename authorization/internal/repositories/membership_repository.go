package repositories

import (
	"authorization/internal/constants"
	"authorization/internal/database"
	"authorization/internal/models"
	"authorization/internal/transport/rest/forms"
	"errors"
	"github.com/google/uuid"
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

func (self *MembershipRepository) InviteMembers(members []forms.MemberInviteForm, roles []string, organizationId uint) error {
	lim := 100
	var chunk []inviteMemberPayloadDto
	var _roles []models.Role
	res := self.db.Where("slug IN ?", roles).Find(&_roles)
	if res.Error != nil {
		return res.Error
	}
	for idx, member := range members {
		var memberRoleId uint = 0
		for _, role := range _roles {
			if role.Slug == member.Role {
				memberRoleId = role.Id
			}
		}
		if memberRoleId == 0 {
			continue
		}
		chunk = append(chunk, inviteMemberPayloadDto{email: member.Email, role: memberRoleId})
		if idx%lim == 0 {
			self.inviteMemberFromMail(chunk, organizationId)
			chunk = make([]inviteMemberPayloadDto, 1)
		}
	}

	return nil
}

type inviteMemberPayloadDto struct {
	email string
	role  uint
}

func (self *MembershipRepository) inviteMemberFromMail(members []inviteMemberPayloadDto, orgId uint) error {
	var emails []string
	userMailRole := map[string]uint{}
	for _, member := range members {
		emails = append(emails, member.email)
		userMailRole[member.email] = member.role
	}
	var users []models.User
	res := self.db.Where("email IN ?", emails).Find(&users)
	if res.Error != nil {
		return res.Error
	}
	var invite []models.OrganizationInvite
	for _, user := range users {
		url, _ := uuid.NewRandom()
		invite = append(invite, models.OrganizationInvite{
			UserId:         user.Id,
			OrganizationId: orgId,
			RoleId:         userMailRole[user.Email],
			Link:           url.String(),
			Status:         int(models.ORG_INIVITED),
		})
	}

	res = self.db.Create(&invite)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
