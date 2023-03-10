package models

type BaseUser struct {
	BaseModel
	Email    string  `gorm:"unique" json:"email"`
	Phone    string  `gorm:"unique" json:"phone"`
	FullName string  `json:"full_name"`
	Avatar   *string `json:"avatar"`
	Bio      *string `json:"bio"`
	Active   bool    `json:"active"`
	RoleId   *uint   `json:"role_id"`
}

type User struct {
	BaseModel
	BaseUser
	Password          string       `json:"password"`
	PasswordResetCode *string      `json:"password_reset_code"`
	Organization      Organization `gorm:"many2many:memberships"`
}

type PublicUser struct {
	BaseModel
	BaseUser
}
