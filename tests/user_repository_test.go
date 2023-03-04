package tests

import (
	"edugree_auth/internal/database"
	users "edugree_auth/internal/repositories"
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
	rep := users.NewRepository()
	var id uint = 2
	model := rep.GetUserById(id)
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
	mock := users.DataPayload{
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
	rep := users.NewRepository()
	err, model := rep.CreateNewUser(&mock)
	if err != nil {
		t.Fatal("User wasnt created", err.Error())
	} else {
		t.Log("User created")
		t.Log(model)
	}
}

func TestRepository_UpdateUser(t *testing.T) {
	mock := users.DataPayload{
		Email:             "test2@mail.com",
		PasswordResetCode: nil,
		Phone:             "+72822922020",
		FullName:          "Test User Com",
		Avatar:            nil,
		Bio:               nil,
		Active:            false,
		RoleId:            nil,
	}
	rep := users.NewRepository()
	var id uint = 3
	err, model := rep.UpdateUser(id, &mock)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}

func TestRepository_DeleteUser(t *testing.T) {
	rep := users.NewRepository()
	var id uint = 3
	err, success := rep.DeleteUser(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(success)
}
