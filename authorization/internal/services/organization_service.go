package services

import (
	casbin2 "authorization/internal/casbin"
	"authorization/internal/helpers"
	"authorization/internal/models"
	"authorization/internal/transport/rest/forms"
)

type OrganizationRepository interface {
	CreateOrganization(organization *models.Organization) (*models.Organization, error)
	GetOrganization(id uint) (*models.Organization, error)
	CheckOrganization(email string, domain string) (*models.Organization, error)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
	AttachOrganizationOwner(orgId uint, userId uint) error
}

type OrganizationService struct {
	repository OrganizationRepository
}

func NewOrganizationService(repository OrganizationRepository) OrganizationService {
	return OrganizationService{
		repository: repository,
	}
}

type OrganizationExistError struct {
	error string
}

func (err *OrganizationExistError) Error() string {
	return "organization is exist"
}

func (self *OrganizationService) CreateOrganization(form *forms.OrganizationCreate, userId uint) (*models.Organization, error) {
	org, err := self.repository.CheckOrganization(form.Email, form.Domain)
	if org != nil && org.Id > 0 {
		return nil, helpers.NewError("organization exist")
	}

	var organization models.Organization
	organization.Title = form.Title
	organization.Email = form.Email
	organization.Domain = form.Domain
	organization.Description = &form.Description
	_, err = self.repository.CreateOrganization(&organization)

	if err != nil {
		return nil, err
	}

	// Добавляем организацию в касбин
	err = casbin2.DefineInitialPolicies(form.Domain)
	if err != nil {
		return nil, err
	}

	err = self.repository.AttachOrganizationOwner(organization.Id, userId)

	return &organization, nil
}

func (self *OrganizationService) GetOrganization(id uint) (*models.Organization, error) {
	organization, err := self.repository.GetOrganization(id)
	if err != nil {
		return nil, err
	}

	return organization, nil
}
