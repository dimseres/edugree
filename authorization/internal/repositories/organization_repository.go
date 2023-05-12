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
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type OrganizationRepository struct {
	BaseRepositoryHelpers
}

func NewOrganizationRepository(requestUuid string) OrganizationRepository {
	return OrganizationRepository{
		BaseRepositoryHelpers{
			db:          database.GetConnection(),
			requestUuid: requestUuid,
		},
	}
}

func (self *OrganizationRepository) CreateOrganization(organization *models.Organization, ownerId uint) (*models.Organization, error) {
	res := self.db.Create(&organization)
	if res.Error != nil {
		return nil, res.Error
	}

	err := self.CreateTenantOrganization(organization, ownerId)
	if err != nil {
		return nil, err
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

	resp := self.db.First(&org, "id = ?", id)
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

func (self *OrganizationRepository) GetOrganizationWithMembers(orgId uint) (*models.Organization, error) {
	var org models.Organization
	resp := self.db.Preload("Members").Preload("Members.User").Preload("Members.Role").First(&org, orgId)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &org, nil
}

func (self *OrganizationRepository) CreateTenantOrganization(organization *models.Organization, userId uint) error {
	var user models.User
	self.db.Where("id = ?", userId).First(&user)

	orgData := map[string]interface{}{
		"id":          organization.Id,
		"name":        organization.Title,
		"domain":      organization.Domain,
		"tenant_uuid": organization.TenantUuid,
		"user": map[string]interface{}{
			"id":    user.Id,
			"name":  user.FullName,
			"email": user.Email,
			"phone": user.Phone,
		},
	}

	postBody, _ := json.Marshal(orgData)
	responseBody := bytes.NewBuffer(postBody)

	url := os.Getenv("COURSE_URL") + "/integration/organization/create"
	if self.requestUuid == "" {
		return errors.New("empty request-id")
	}

	req, err := http.NewRequest(http.MethodPost, url, responseBody)
	if err != nil {
		return err
	}

	req.Header.Set("X-ACCESS-KEY", os.Getenv("GATEWAY_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-REQUEST-ID", self.requestUuid)

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	responseStructure := struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}{}
	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		json.Unmarshal(body, &responseStructure)

		if err != nil {
			return err
		}

		if responseStructure.Error && responseStructure.Message != "" {
			return errors.New(responseStructure.Message)
		}

		return errors.New(resp.Status)
	}

	return nil
}
