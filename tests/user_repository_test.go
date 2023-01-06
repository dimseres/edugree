package tests

import (
	"edugree_auth/internal/database"
	"edugree_auth/internal/repositories/users"
	"github.com/joho/godotenv"
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
	id := 2
	model := rep.GetUserById(id)
	if model.Id == 0 {
		t.Fatalf("User With ID: %d not found", id)
	}
	t.Log("User Found")
	t.Log(model)
}
