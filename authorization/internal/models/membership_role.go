package models

type MembershipRole struct {
	MembershipId uint `gorm:"primaryKey;autoIncrement:false"`
	UserId       uint `gorm:"primaryKey;autoIncrement:false"`
	RoleId       uint `gorm:"primaryKey;autoIncrement:false"`
	User         *User
	Membership   *Membership
	Role         *Role `gorm:"foreignKey:role_id"`
}
