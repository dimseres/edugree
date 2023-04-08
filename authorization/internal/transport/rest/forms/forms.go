package forms

type UserRegistrationForm struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Phone    string `form:"phone" json:"phone" validate:"required,e164"`
	Password string `form:"password" json:"password" validate:"required"`
	FullName string `form:"full_name" json:"full_name" validate:"required"`
}

type UserLoginForm struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
}

type SetTenantForm struct {
	TenantId uint `form:"tenant_id" json:"tenant_id" validate:"required"`
}

type OrganizationCreate struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Domain      string `json:"domain" form:"domain" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Description string `json:"description" form:"title"`
}
