package services

import (
	"authorization/internal/models"
)

type MembershipRepository interface {
	GetMembershipData(organizationId uint, userId uint) (*models.Membership, error)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
	DeleteMember(memberId uint, organizationId uint) error
}

type MembershipService struct {
	repository MembershipRepository
	BaseService
}

func NewMembershipService(repository MembershipRepository, ctx *TenantContext) MembershipService {
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
