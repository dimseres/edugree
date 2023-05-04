package models

type Membership struct {
	UserId         uint          `gorm:"foreignKey;autoIncrement:false" json:"user_id"`
	OrganizationId uint          `gorm:"foreignKey;autoIncrement:false" json:"organization_id"`
	RoleId         uint          `gorm:"foreignKey;autoIncrement:false" json:"role_id"`
	Organization   *Organization `json:"organization"`
	Role           *Role         `json:"role"`
	User           User          `json:"user"`
}
