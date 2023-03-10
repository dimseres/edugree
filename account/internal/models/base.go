package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uint           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
