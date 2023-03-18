package tests

import (
	"authorization/internal/database"
	users "authorization/internal/repositories"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	database.InitConnection()
}

func TestRepository_GetUserById(t *testing.T) {
	rep := users.NewUserRepository()
	var id uint = 4
	model, err := rep.GetUserById(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
	if model.BaseModel.Id == 0 {
		t.Fatalf("User With ID: %d not found", id)
	}
	t.Log("User Found")
	t.Log(model)
}

func TestRepository_CreateNewUser(t *testing.T) {
	hashed, err := bcrypt.GenerateFromPassword([]byte("awda"), 11)
	if err != nil {
		panic(err)
	}
	mock := users.UserDataPayload{
		Email:             "test@mail.com",
		Password:          string(hashed),
		PasswordResetCode: nil,
		Phone:             "+72822922020",
		FullName:          "Test User Com",
		Avatar:            nil,
		Bio:               nil,
		Active:            false,
		RoleId:            nil,
	}
	rep := users.NewUserRepository()
	err, model := rep.CreateNewUser(&mock)
	if err != nil {
		t.Fatal("User wasnt created", err.Error())
	} else {
		t.Log("User created")
		t.Log(model)
	}
}

func TestRepository_UpdateUser(t *testing.T) {
	mock := users.UserDataPayload{
		Email:             "test2@mail.com",
		PasswordResetCode: nil,
		Phone:             "+72822922020",
		FullName:          "Test User Com",
		Avatar:            nil,
		Bio:               nil,
		Active:            false,
		RoleId:            nil,
	}
	rep := users.NewUserRepository()
	var id uint = 3
	err, model := rep.UpdateUser(id, &mock)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}

func TestRepository_DeleteUser(t *testing.T) {
	rep := users.NewUserRepository()
	var id uint = 3
	err, success := rep.DeleteUser(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(success)
}
