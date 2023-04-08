package repositories

// Email             string
//	Password          string
//	PasswordResetCode *string
//	Phone             string
//	FullName          string
//	Avatar            *string
//	Bio               *string
//	Active            bool

import (
	"authorization/internal/database"
	"authorization/internal/models"
)

type OrganizationRepository struct {
	BaseRepositoryHelpers
}

func NewOrganizationRepository() OrganizationRepository {
	return OrganizationRepository{
		BaseRepositoryHelpers{
			db: database.InitConnection(),
		},
	}
}

func (self *OrganizationRepository) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	res := self.db.Create(&organization)
	if res.Error != nil {
		return nil, res.Error
	}

	return organization, nil
}

func (self *OrganizationRepository) CheckOrganization(email string, domain string) (*models.Organization, error) {
	var organization models.Organization

	response := self.db.Where("email = ? or domain = ?", email, domain).First(&organization)
	if response.Error != nil {
		return nil, response.Error
	}

	return &organization, nil
}

func (self *OrganizationRepository) GetOrganization(id uint) (*models.Organization, error) {
	var org models.Organization

	resp := self.db.Find(&org, id)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &org, nil
}

func (self *OrganizationRepository) AttachOrganizationOwner(orgId uint, userId uint) error {
	var member models.Membership
	var role models.Role
	res := self.db.Where("slug = ?", "owner").First(&role)
	if res.Error != nil {
		return res.Error
	}
	res = self.db.FirstOrCreate(&member, models.Membership{
		UserId:         userId,
		OrganizationId: orgId,
		RoleId:         role.Id,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
