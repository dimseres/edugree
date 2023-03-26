package services

import (
	"authorization/internal/models"
)

type MembershipRepository interface {
	GetMembershipData(organizationId uint, userId uint) (*models.Membership, error)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
}

type MembershipService struct {
	repository MembershipRepository
}

func NewMembershipService(repository MembershipRepository) MembershipService {
	return MembershipService{
		repository: repository,
	}
}

func (self *MembershipService) GetMembershipData(orgId uint, userId uint) (*models.Membership, error) {
	member, err := self.repository.GetMembershipData(orgId, userId)
	return member, err
}
