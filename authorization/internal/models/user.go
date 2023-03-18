package models

import (
	"gorm.io/gorm"
	"time"
)

type ModelTime struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type BaseModel struct {
	Id uint `gorm:"primaryKey;autoIncrement" json:"id"`
}

//type BaseUser struct {
//	BaseModel
//	Email    string  `gorm:"unique" json:"email"`
//	Phone    string  `gorm:"unique" json:"phone"`
//	FullName string  `json:"full_name"`
//	Avatar   *string `json:"avatar"`
//	Bio      *string `json:"bio"`
//	Active   bool    `json:"active"`
//}

type User struct {
	BaseModel

	Email             string         `gorm:"not null;unique" json:"email"`
	Password          string         `gorm:"not null" json:"-"`
	PasswordResetCode *string        `gorm:"size:256" json:"-"`
	Phone             string         `gorm:"size:256;not null;unique" json:"phone"`
	FullName          string         `gorm:"size:512;not null" json:"full_name"`
	Avatar            *string        `json:"avatar"`
	Bio               *string        `gorm:"type:text" json:"bio"`
	Active            bool           `gorm:"not null;default:true" json:"active"`
	Membership        []Organization `gorm:"many2many:memberships"`
	DomainRole        []Role         `gorm:"many2many:memberships"`

	ModelTime
}

//type PublicUser struct {
//	Role  *[]Role  `gorm:"foreignKey:role_id;references:id" json:"role"`
//	Token *[]Token `json:"token"`
//}
