package models

type Organization struct {
	BaseModel
	Title       string
	Email       string
	Description *string
	Avatar      *string
	Domain      string
	Members     *[]Members `gorm:"many2many:memberships"`
}

type Members struct {
	UserId         uint
	OrganizationId uint
}
