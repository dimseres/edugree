package models

type Service struct {
	BaseModel

	Title       string          `gorm:"size:256;index;not null" json:"title"`
	Slug        string          `gorm:"unique;not null" json:"slug"`
	Description *string         `gorm:"type:text"`
	User        *[]Organization `gorm:"many2many:organizations_services"`

	ModelTime
}
