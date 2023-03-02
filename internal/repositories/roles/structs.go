package users

import (
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

type RoleDataPayload struct {
	Title  string
	Domain string
	Slug   string
}
