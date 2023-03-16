package models

type Role struct {
	BaseModel

	Name           string          `gorm:"not null;index;uniqueIndex:unique_index" json:"title"`
	Slug           string          `gorm:"not null;index;unique;uniqueIndex:unique_index" json:"slug"`
	Description    *string         `gorm:"type:text" json:"description"`
	IsSystem       bool            `gorm:"default:false" json:"is_system"`
	OrganizationId *uint           `gorm:"uniqueIndex:unique_index"`
	Organizations  *[]Organization `gorm:"many2many:memberships" json:"organization"`
	Users          *[]User         `gorm:"many2many:memberships" json:"user"`
	//Permissions    *[]Permissions  `gorm:"many2many:organizations;foreignKey:name;joinForeignKey:domain;References:v0;joinReferences:v3"`

	ModelTime
}
