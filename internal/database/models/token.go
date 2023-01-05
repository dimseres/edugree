package models

import (
	"edugree_auth/internal/database"
	"time"
)

type Token struct {
	database.BaseModel
	User       User `gorm:"foreignKey:user_id;references:id"`
	UserId     uint
	token      string
	ExpireDate time.Time
}
