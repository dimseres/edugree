package models

type InviteStatus int

const (
	ORG_INIVITED InviteStatus = iota + 1
	ORG_REJECTED
	ORG_ACCEPTED
)

func (s InviteStatus) String() string {
	switch s {
	case ORG_INIVITED:
		return "invited"
	case ORG_REJECTED:
		return "rejected"
	case ORG_ACCEPTED:
		return "accepted"
	}
	return "unknown"
}

type OrganizationInvite struct {
	BaseModel

	OrganizationId uint   `json:"organization_id"`
	UserId         uint   `json:"user_id"`
	RoleId         uint   `json:"role_id"`
	Link           string `json:"link"`
	Status         int    `json:"status"`

	User         User         `json:"user"`
	Organization Organization `json:"organization"`
	Role         Role         `json:"role"`
	//Permissions *[]Permissions `gorm:"foreignKey:v3"`

	ModelTime
}
