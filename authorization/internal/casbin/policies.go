package casbin

const SubOwner = "owner"
const SubAdmin = "administrator"
const SubModer = "moderator"
const SubTeacher = "teacher"
const SubStudent = "student"

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
	SubOwner,
	SubAdmin,
	SubModer,
	SubTeacher,
	SubStudent,
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
	SubOwner: {
		ObjUsers:    {ActCreate, ActRead, ActUpdate, ActDelete, ActModifyAdmin},
		ObjRoles:    {ActCreate, ActRead, ActUpdate, ActDelete},
		ObjCourses:  {ActCreate, ActRead, ActUpdate, ActDelete},
		ObjServices: {ActCreate, ActRead, ActUpdate, ActDelete},
	},
	SubAdmin: {
		ObjUsers:   {ActCreate, ActRead, ActUpdate, ActDelete},
		ObjCourses: {ActCreate, ActRead, ActUpdate, ActDelete},
	},
	SubModer: {
		ObjCourses: {ActCreate, ActRead, ActUpdate, ActDelete},
	},
	SubStudent: {
		ObjCourses: {ActRead},
	},
}
