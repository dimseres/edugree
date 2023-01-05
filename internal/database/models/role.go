package models

import "edugree_auth/internal/database"

type Role struct {
	database.BaseModel
	Title  string
	Domain string
	Slug   string `gorm:"unique"`
	User   *[]User
}
