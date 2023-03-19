package tests

import (
	"authorization/internal/database"
	users "authorization/internal/repositories"
	"authorization/internal/services"
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

func TestService_SignIn(t *testing.T) {
	rep := users.NewAuthRepository()
	service := services.NewAuthService(&rep)

	err, result := service.SignIn("test1@example.net", "admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestService_SignInJwt(t *testing.T) {
	rep := users.NewAuthRepository()
	service := services.NewAuthService(&rep)
	err, result := service.SignIn("test1@example.net", "admin")

	if err != nil {
		t.Fatal(err)
	}

	err, jwt := service.CreateJwtToken(result, "domain")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(jwt)
}
