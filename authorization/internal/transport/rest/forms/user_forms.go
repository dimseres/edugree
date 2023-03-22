package forms

type UserRegistrationForm struct {
	Email    string `form:"email" validate:"required,email"`
	Phone    string `form:"phone" validate:"required,e164"`
	Password string `form:"password" validate:"required"`
	FullName string `form:"full_name" validate:"required"`
}
