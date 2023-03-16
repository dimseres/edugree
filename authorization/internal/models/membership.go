package models

type Membership struct {
	UserId         uint          `gorm:"foreignKey;autoIncrement:false"`
	OrganizationId uint          `gorm:"foreignKey;autoIncrement:false"`
	RoleId         uint          `gorm:"foreignKey;autoIncrement:false"`
	Organization   *Organization `gorm:"foreignKey:organization_id"`
	Role           *Role         `gorm:"foreignKey:role_id"`
	User           *User         `gorm:"foreignKey:user_id"`
}
