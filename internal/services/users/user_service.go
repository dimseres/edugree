package users

import (
	"edugree_auth/internal/repositories/users"
)

func NewSerivice() UserService {
	return UserService{
		repository: users.NewRepository(),
	}
}

func GetUserById(id int) {

}
