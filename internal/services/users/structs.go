package users

import (
	"edugree_auth/internal/repositories/users"
)

type UserService struct {
	repository users.Repository
}
