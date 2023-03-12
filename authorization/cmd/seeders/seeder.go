package main

import (
	"authorization/internal/database"
	"authorization/internal/helpers"
	"authorization/internal/models"
	"authorization/internal/repositories"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	connection := database.InitConnection()
	createUsers(connection)
}

func createRoles(db *gorm.DB) {
	roles := []models.Role{}
	db.Create()
}

func createOrganization(db *gorm.DB) {
	roles := models.Organization{
		BaseUser:          models.BaseUser{},
		Password:          "",
		PasswordResetCode: nil,
		Role:              nil,
		Token:             nil,
	}
	db.Create()
}

func createUsers(db *gorm.DB) {
	userRepo := repositories.NewUserRepository()
	hashPassword, err := helpers.CreatePasswordHash("admin")
	if err != nil {
		panic(err)
	}

	err, _ = userRepo.CreateNewUser(&repositories.UserDataPayload{
		Email:             "test1@example.net",
		Password:          hashPassword,
		PasswordResetCode: nil,
		Phone:             "+111",
		FullName:          "Test User One",
		Avatar:            nil,
		Bio:               nil,
		Active:            false,
		RoleId:            nil,
	})
	if err != nil {
		panic(err)
	}
}
