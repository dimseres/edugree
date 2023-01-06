package models

import "edugree_auth/internal/database"

type User struct {
	database.BaseModel
	Email             string `gorm:"unique"`
	Password          string
	PasswordResetCode *string
	Phone             string `gorm:"unique"`
	FullName          string
	Avatar            *string
	Bio               *string
	Active            bool
	RoleId            *uint
	Role              *Role `gorm:"foreignKey:role_id;references:id"`
	Token             *[]Token
}
