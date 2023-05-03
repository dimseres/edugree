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

	organization_id uint64
	user_id         uint64
	link            string
	status          int

	Users        User
	Organization Organization
	//Permissions *[]Permissions `gorm:"foreignKey:v3"`

	ModelTime
}
