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

	OrganizationId uint
	UserId         uint
	RoleId         uint
	Link           string
	Status         int

	User         User
	Organization Organization
	Role         Role
	//Permissions *[]Permissions `gorm:"foreignKey:v3"`

	ModelTime
}
