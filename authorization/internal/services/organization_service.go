package services

import (
	casbin2 "authorization/internal/casbin"
	"authorization/internal/helpers"
	"authorization/internal/models"
	"authorization/internal/transport/rest/forms"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(organization *models.Organization, ownerID uint) (*models.Organization, error)
	GetOrganization(id uint) (*models.Organization, error)
	CheckOrganization(email string, domain string) (*models.Organization, error)
	LoadRelation(model interface{}, relation ...string) (interface{}, error)
	AttachOrganizationOwner(orgId uint, userId uint) error
	StartTransaction()
	EndTransaction()
	RollbackTransaction()
	GetDb() *gorm.DB
}

type OrganizationService struct {
	repository OrganizationRepository
	BaseService
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
	_uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	organization.Title = form.Title
	organization.Email = form.Email
	organization.Domain = form.Domain
	organization.Description = &form.Description
	organization.TenantUuid = _uuid.String()

	self.repository.StartTransaction()
	_, err = self.repository.CreateOrganization(&organization, userId)

	if err != nil {
		self.repository.RollbackTransaction()
		return nil, err
	}

	// Добавляем организацию в касбин
	err = casbin2.DefineInitialPolicies(form.Domain)
	if err != nil {
		self.repository.RollbackTransaction()
		return nil, err
	}

	err = self.repository.AttachOrganizationOwner(organization.Id, userId)
	if err != nil {
		self.repository.RollbackTransaction()
		return nil, err
	}

	self.repository.EndTransaction()
	return &organization, nil
}

func (self *OrganizationService) GetOrganization(id uint) (*models.Organization, error) {
	organization, err := self.repository.GetOrganization(id)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

//func (self *OrganizationService) CreateCourseDomain(organization *models.Organization, userId uint) error {
//	var user models.User
//	self.repository.GetDb().Where("id = ?", userId).First(&user)
//
//	orgData := map[string]interface{}{
//		"id":          organization.Id,
//		"name":        organization.Title,
//		"domain":      organization.Domain,
//		"tenant_uuid": organization.TenantUuid,
//		"owner": map[string]interface{}{
//			"id":    user.Id,
//			"name":  user.FullName,
//			"email": user.Email,
//			"phone": user.Phone,
//		},
//	}
//
//	postBody, _ := json.Marshal(orgData)
//	responseBody := bytes.NewBuffer(postBody)
//	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
//	defer resp.Body.Close()
//	if err != nil {
//		return err
//	}
//
//	responseStructure := struct {
//		Error   bool   `json:"error"`
//		Message string `json:"message"`
//	}{}
//	body, err := io.ReadAll(resp.Body)
//	json.Unmarshal(body, &responseStructure)
//
//	if responseStructure.Error {
//		return errors.New(responseStructure.Message)
//	}
//
//	return nil
//}
