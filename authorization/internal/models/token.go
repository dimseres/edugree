package models

import (
	"time"
)

type Token struct {
	BaseModel

	UserId   uint      `json:"user_id"`
	Token    string    `json:"token"`
	Salt     string    `json:"-"`
	ExpireAt time.Time `json:"expire_at"`
	User     User      `gorm:"foreignKey:user_id;references:id" json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
