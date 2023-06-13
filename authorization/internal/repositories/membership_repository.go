package repositories

import (
	"authorization/config"
	"authorization/internal/constants"
	"authorization/internal/database"
	"authorization/internal/models"
	"authorization/internal/transport/rest/forms"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

type MembershipRepository struct {
	BaseRepositoryHelpers
}

func NewMembershipRepository(requestUuid string) MembershipRepository {
	return MembershipRepository{
		BaseRepositoryHelpers{
			db:          database.GetConnection(),
			requestUuid: requestUuid,
		},
	}
}

func (self *MembershipRepository) GetMembershipData(userId uint, organizationId uint) (*models.Membership, error) {
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

func (self *MembershipRepository) ChangeMemberRole(userId uint, roleId string, organizationId uint) error {
	var role models.Role
	res := self.db.Where("user_id = ? AND role_id = ?", userId, roleId).Find(&role)
	if res.Error != nil {
		return res.Error
	}

	res = self.db.Model(&role).Update("roleId", role.Id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (self *MembershipRepository) InviteMembers(members []forms.MemberInviteForm, roles []string, organizationId uint) error {
	lim := 100
	var chunk []inviteMemberPayloadDto
	var _roles []models.Role
	//var organization models.Organization
	//res := self.db.Find(&organization, organizationId)
	//if res.Error != nil {
	//	return res.Error
	//}

	var membership []models.Membership
	res := self.db.Joins("User").Where("organization_id = ?", organizationId).Find(&membership)
	if res.Error != nil {
		return res.Error
	}

	membersSet := make(map[string]bool)
	for _, member := range membership {
		membersSet[member.User.Email] = true
	}

	res = self.db.Where("slug IN ?", roles).Find(&_roles)

	if res.Error != nil {
		return res.Error
	}
	for idx, member := range members {
		inList := membersSet[member.Email]
		if inList {
			continue
		}
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
	var userIds []uint
	for _, user := range users {
		url, _ := uuid.NewRandom()
		invite = append(invite, models.OrganizationInvite{
			UserId:         user.Id,
			OrganizationId: orgId,
			RoleId:         userMailRole[user.Email],
			Link:           url.String(),
			Status:         int(models.ORG_INIVITED),
		})
		userIds = append(userIds, user.Id)
	}

	var invited []models.OrganizationInvite
	res = self.db.Where("user_id IN ? and organization_id = ? and status = ?", userIds, orgId, models.ORG_ACCEPTED).Find(&invited)
	if res.Error != nil {
		return res.Error
	}

	var filteredInvites []models.OrganizationInvite
	if invited != nil && len(invited) > 0 {
		for _, i := range invite {
			for _, j := range invited {
				if i.UserId != j.UserId {
					filteredInvites = append(filteredInvites, j)
				}
			}
		}
	} else {
		filteredInvites = invite
	}

	res = self.db.Create(&filteredInvites)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (self *MembershipRepository) GetInviteList(page int, perPage int, orgId uint) (*[]models.OrganizationInvite, error) {
	var invites []models.OrganizationInvite
	pagination := PaginationConfig{
		Page:    page,
		PerPage: perPage,
	}
	var total int64
	res := self.db.Scopes(self.Paginate(&pagination)).Preload("User").Preload("Role").Where("organization_id = ?", orgId).Find(&invites).Count(&total)
	if res.Error != nil {
		return nil, res.Error
	}

	return &invites, nil
}

func (self *MembershipRepository) AddUserMemberShip(userId uint, organizationId uint, roleId uint) (*models.Membership, error) {
	membership := models.Membership{
		UserId:         userId,
		OrganizationId: organizationId,
		RoleId:         roleId,
	}
	res := self.db.Create(&membership)
	if res.Error != nil {
		return nil, res.Error
	}
	return &membership, nil
}

func (self *MembershipRepository) RejectOrAcceptInvite(userId uint, link string, action string) (*models.OrganizationInvite, error) {
	status := 0
	newMember := false
	switch action {
	case models.ORG_ACCEPTED.String():
		status = int(models.ORG_ACCEPTED)
		newMember = true
	case models.ORG_REJECTED.String():
		status = int(models.ORG_REJECTED)
	}

	var invite models.OrganizationInvite
	res := self.db.Where("link = ? AND user_id = ? AND status = ?", link, userId, models.ORG_INIVITED).Joins("User").Joins("Organization").Joins("Role").Find(&invite)
	if res.Error != nil || invite.Id == 0 {
		if invite.Id == 0 {
			return nil, errors.New("link not found")
		}
		return nil, res.Error
	}

	res = self.db.Model(&invite).Update("status", status)
	if res.Error != nil {
		return nil, res.Error
	}

	if newMember {
		_, err := self.AddUserMemberShip(invite.UserId, invite.OrganizationId, invite.RoleId)
		if err != nil {
			return nil, err
		}

		err = self.CreateTenantMember(&invite.Organization, &invite.User, invite.Role.Slug)
		if err != nil {
			return nil, err
		}

		// Помечаем предыдущие приглашения как отклоненные
		res = self.db.Model(models.OrganizationInvite{}).Where("link <> ?", link).Delete(models.OrganizationInvite{
			Status: int(models.ORG_REJECTED),
		})

		if res.Error != nil {
			return nil, err
		}
	}

	return &invite, nil
}

func (self *MembershipRepository) CreateTenantMember(organization *models.Organization, user *models.User, role string) error {
	orgData := map[string]interface{}{
		"domain":      organization.Domain,
		"tenant_uuid": organization.TenantUuid,
		"user": map[string]interface{}{
			"id":    user.Id,
			"name":  user.FullName,
			"email": user.Email,
			"phone": user.Phone,
			"role":  role,
		},
	}

	postBody, _ := json.Marshal(orgData)
	responseBody := bytes.NewBuffer(postBody)

	url := config.GetConfig("COURSE_URL") + "/integration/organization/member/add"
	if self.requestUuid == "" {
		return errors.New("empty request-id")
	}

	req, err := http.NewRequest(http.MethodPost, url, responseBody)
	if err != nil {
		return err
	}

	req.Header.Set("X-ACCESS-KEY", config.GetConfig("GATEWAY_KEY"))
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
