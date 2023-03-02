package users

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserDataPayload struct {
	Email             string
	Password          string
	PasswordResetCode *string
	Phone             string
	FullName          string
	Avatar            *string
	Bio               *string
	Active            bool
	RoleId            *uint
}
