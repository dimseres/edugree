package tests

import (
	"edugree_auth/internal/database"
	users "edugree_auth/internal/repositories"
	"edugree_auth/internal/services"
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

func TestService_SignIn(t *testing.T) {
	rep := users.NewAuthRepository()
	service := services.NewAuthService(&rep)

	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

	fmt.Println(string(password))

	if err != nil {
		panic(err)
	}

	err, result := service.SignIn("admin@example.com", "admin")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
