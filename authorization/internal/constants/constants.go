package constants

import "time"

const SubOwner = "owner"
const SubAdmin = "administrator"
const SubModer = "moderator"
const SubTeacher = "teacher"
const SubStudent = "student"

const REFRESH_LIFETIME = time.Hour * 24 * 7
const JWT_LIFETIME = time.Second * 900 // JWT_LIFETIME lifetime of jwt token

// SAGA Events name below

const CREATE_ORGANIZATION_EVENT = "create_organization"
const ROLLBACK_ORGANIZATION_EVENT = "rollback_organization_create"
const CREATE_USERS = "create_users"
const ROLLBACK_USERS = "rollback_users_create"
const CREATE_USER = "create_user"
const ROLLBACK_USER = "rollback_user_create"
const UPDATE_USER = "update_user"
const ROLLBACK_USER_UPDATE = "rollback_user_update"
const DELETE_USER = "delete_user"
const ROLLBACK_USER_DELETE = "rollback_user_delete"

// end of events
