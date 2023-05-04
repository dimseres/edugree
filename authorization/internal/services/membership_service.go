package services

import (
	"authorization/internal/helpers"
	"authorization/internal/models"
	"authorization/internal/transport/rest/forms"
	"errors"
)

type MembershipRepository interface {
	GetMembershipData(organizationId uint, userId uint) (*models.Membership, error)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
	DeleteMember(memberId uint, organizationId uint) error
	InviteMembers(members []forms.MemberInviteForm, organizationId uint) error
}

type MembershipService struct {
	repository MembershipRepository
	BaseService
}

func NewMembershipService(repository MembershipRepository, ctx *helpers.TenantContext) MembershipService {
	return MembershipService{
		repository:  repository,
		BaseService: BaseService{tenantContext: ctx},
	}
}

func (self *MembershipService) GetMembershipData(orgId uint, userId uint) (*models.Membership, error) {
	member, err := self.repository.GetMembershipData(orgId, userId)
	return member, err
}

func (self *MembershipService) DeleteMember(userId uint) error {
	return self.repository.DeleteMember(userId, self.tenantContext.Id)
}

func (self *MembershipService) InviteMembers(form *forms.InviteMembersForm) (bool, error) {
	eventorRole := self.BaseService.tenantContext.Role
	availableRoles, ok := helpers.GetCreateAvailableRoles(eventorRole)

	if !ok {
		return false, errors.New("wrong user role")
	}

	for _, newMember := range form.Members {
		canCreate := false
		for _, available := range availableRoles {
			if newMember.Role == available {
				canCreate = true
				break
			}
		}
		if !canCreate {
			return false, errors.New("Not enough permission for create role")
		}
	}

	err := self.repository.InviteMembers(form.Members, self.tenantContext.Id)
	if err != nil {
		return false, err
	}

	return false, nil
}
