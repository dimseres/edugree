package responses

import "authorization/internal/models"

type ServiceResponse struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type OrganizationResponse struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Domain string `json:"domain"`
	//Email       string             `json:"email"`
	//Description *string            `json:"full_name"`
	Avatar *string `json:"avatar"`
	//Bio         *string            `json:"bio"`
	//Active      bool               `json:"active"`
	//Services *[]ServiceResponse `json:"services"`
}

type MembershipResponse struct {
	//UserId         uint                  `json:"user_id"`
	//OrganizationId uint                  `json:"organization_id"`
	//RoleId         uint                  `json:"role_id"`
	Organization *OrganizationResponse `json:"organization"`
	Role         *RoleResponse         `json:"role"`
}

type RoleResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"title"`
	Slug string `json:"slug"`
}

type UserLoginResponse struct {
	Id         uint                 `json:"id"`
	Email      string               `json:"email"`
	Phone      string               `json:"phone"`
	FullName   string               `json:"full_name"`
	Avatar     *string              `json:"avatar"`
	Bio        *string              `json:"bio"`
	Active     bool                 `json:"active"`
	Membership []MembershipResponse `json:"membership"`
}

func NewUserResponse(user *models.User) *UserLoginResponse {
	response := new(UserLoginResponse)

	response.Id = user.Id
	response.FullName = user.FullName
	response.Email = user.Email
	response.Phone = user.Phone
	response.Bio = user.Bio

	if user.Membership != nil {
		for _, member := range *user.Membership {
			_member := new(MembershipResponse)
			if member.Role != nil {
				_member.Role = &RoleResponse{
					Id:   member.Role.Id,
					Name: member.Role.Name,
					Slug: member.Role.Slug,
				}
			}
			if member.Organization != nil {
				_member.Organization = &OrganizationResponse{
					Id:     member.Organization.Id,
					Title:  member.Organization.Title,
					Domain: member.Organization.Domain,
					Avatar: member.Organization.Avatar,
				}
			}
			response.Membership = append(response.Membership, *_member)
		}
	}
	return response
}

type SetTenantResponse struct {
	Role         *RoleResponse        `json:"role"`
	Organization OrganizationResponse `json:"organization"`
	Service      *[]ServiceResponse   `json:"services"`
}

func NewSetTenantResponse(member *models.Membership) *SetTenantResponse {
	response := new(SetTenantResponse)

	response.Organization = OrganizationResponse{
		Id:     member.Organization.Id,
		Title:  member.Organization.Title,
		Domain: member.Organization.Domain,
		Avatar: member.Organization.Avatar,
	}

	if member.Role != nil {
		response.Role = &RoleResponse{
			Id:   member.Role.Id,
			Name: member.Role.Name,
			Slug: member.Role.Slug,
		}
	}

	if member.Organization.Services != nil {
		tempArr := new([]ServiceResponse)
		for _, service := range *member.Organization.Services {
			*tempArr = append(*tempArr, ServiceResponse{
				Id:    service.Id,
				Title: service.Title,
				Slug:  service.Slug,
			})
		}
		response.Service = tempArr
	}

	return response
}
