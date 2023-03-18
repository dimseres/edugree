package models

type Membership struct {
	UserId         uint `gorm:"foreignKey;autoIncrement:false"`
	OrganizationId uint `gorm:"foreignKey;autoIncrement:false"`
	RoleId         uint `gorm:"foreignKey;autoIncrement:false"`
	Organization   *Organization
	Role           *Role
	User           *User
}
