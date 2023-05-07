package models

type Organization struct {
	BaseModel

	Title       string        `gorm:"size:512;index;not null" json:"title"`
	Domain      string        `gorm:"size:512;index;unique;not null" json:"domain"`
	SecretKey   string        `gorm:"size:512;index;not null" json:"secret_key"`
	TenantUuid  string        `gorm:"type:text;not null;unique;" json:"tenant_uuid"`
	Email       string        `gorm:"size:256;not null;unique;index" json:"email"`
	Description *string       `gorm:"type:text" json:"description"`
	Avatar      *string       `json:"avatar"`
	Bio         *string       `gorm:"type:text" json:"bio"`
	Active      bool          `json:"active"`
	Users       *[]User       `gorm:"many2many:memberships" json:"users"`
	Roles       *[]Role       `gorm:"many2many:memberships" json:"roles"`
	Services    *[]Service    `gorm:"many2many:organizations_services" json:"services"`
	Members     *[]Membership `json:"members"`
	//Permissions *[]Permissions `gorm:"foreignKey:v3"`

	ModelTime
}
