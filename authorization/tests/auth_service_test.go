package tests

import (
	"authorization/internal/database"
	"authorization/internal/helpers"
	users "authorization/internal/repositories"
	"authorization/internal/services"
	"fmt"
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

	//firstSalt, err := scrypt.Key([]byte("admin"), []byte(os.Getenv("SALT")), 2048, 4, 2, 32)
	//passwordHash, err := scrypt.Key([]byte("admin"), firstSalt, 16384, 8, 1, 32)
	//fmt.Println("============", string(passwordHash))
	//encoded := hex.EncodeToString(passwordHash)
	//fmt.Println("+++++++++++", encoded)
	//password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

	password, err := helpers.CreatePasswordHash("admin")

	fmt.Println("-------------", password)

	if err != nil {
		panic(err)
	}

	err, result := service.SignIn("admin@example.com", "admin")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
