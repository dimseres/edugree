package casbin

import "authorization/internal/constants"

const ObjUsers = "users"
const ObjRoles = "roles"
const ObjServices = "services"
const ObjCourses = "courses"

const ActCreate = "create"
const ActRead = "read"
const ActUpdate = "update"
const ActDelete = "delete"
const ActModifyAdmin = "modify_admin"

var subjects = []string{
	constants.SubOwner,
	constants.SubAdmin,
	constants.SubModer,
	constants.SubTeacher,
	constants.SubStudent,
}

var objects = []string{
	ObjUsers,
	ObjRoles,
	ObjServices,
	ObjCourses,
}

var actions = []string{
	ActCreate,
	ActRead,
	ActUpdate,
	ActDelete,
}

var rolePolicies = map[string]map[string][]string{
	constants.SubOwner: {
		ObjUsers:    {ActCreate, ActRead, ActUpdate, ActDelete, ActModifyAdmin},
		ObjRoles:    {ActCreate, ActRead, ActUpdate, ActDelete},
		ObjCourses:  {ActCreate, ActRead, ActUpdate, ActDelete},
		ObjServices: {ActCreate, ActRead, ActUpdate, ActDelete},
	},
	constants.SubAdmin: {
		ObjUsers:   {ActCreate, ActRead, ActUpdate, ActDelete},
		ObjCourses: {ActCreate, ActRead, ActUpdate, ActDelete},
	},
	constants.SubModer: {
		ObjCourses: {ActCreate, ActRead, ActUpdate, ActDelete},
	},
	constants.SubStudent: {
		ObjCourses: {ActRead},
	},
}
