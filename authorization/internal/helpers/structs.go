package helpers

import "authorization/internal/constants"

type TenantContext struct {
	Id     uint
	Domain string
	UserId uint
	Role   string
}

var roleCreationMap = map[string][]string{
	constants.SubOwner:   {constants.SubAdmin, constants.SubModer, constants.SubTeacher, constants.SubStudent},
	constants.SubAdmin:   {constants.SubModer, constants.SubTeacher, constants.SubStudent},
	constants.SubModer:   {constants.SubTeacher, constants.SubStudent},
	constants.SubTeacher: {constants.SubStudent},
	constants.SubStudent: {},
}

func GetCreateAvailableRoles(userRole string) ([]string, bool) {
	value, ok := roleCreationMap[userRole]
	return value, ok
}
