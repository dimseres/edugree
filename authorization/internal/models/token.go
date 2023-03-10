package models

import (
	"time"
)

type BaseToken struct {
	UserId     uint      `json:"user_id"`
	Token      string    `json:"token"`
	ExpireDate time.Time `json:"expire_date"`
}

type Token struct {
	BaseModel
	User User `gorm:"foreignKey:user_id;references:id" json:"user"`
	BaseToken
}

type PublicToken struct {
	BaseModel
	BaseToken
}
