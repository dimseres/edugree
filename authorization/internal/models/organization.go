package models

type Organization struct {
	BaseModel

	Title       string     `gorm:"size:512;index;not null" json:"title"`
	Domain      string     `gorm:"size:512;index;unique;not null" json:"domain"`
	Email       string     `gorm:"size:256;not null;unique;index" json:"email"`
	Description *string    `gorm:"type:text" json:"full_name"`
	Avatar      *string    `json:"avatar"`
	Bio         *string    `gorm:"type:text" json:"bio"`
	Active      bool       `json:"active"`
	User        *[]User    `gorm:"many2many:memberships"`
	Roles       *[]Role    `gorm:"many2many:roles"`
	Services    *[]Service `gorm:"many2many:organizations_services"`
	//Permissions *[]Permissions `gorm:"foreignKey:v3"`

	ModelTime
}
