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
