package main

import (
	"authorization/internal/casbin"
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
	_ = createRoles(connection)
	user := createUsers(connection)
	organization := createOrganization(connection, *user)

	casbin.DefineInitialPolicies(organization.Domain)
}

func createRoles(db *gorm.DB) *[]models.Role {
	orgDesc := "Владелец организации"
	adminDesc := "Администратор организации"
	moderatorDesc := "Модератор организации"
	teacherDesc := "Преподаватель организации"
	studentDesc := "Студент организации"

	owner := models.Role{
		Name:           "Владелец",
		Description:    &orgDesc,
		Slug:           casbin.SubOwner,
		IsSystem:       true,
		OrganizationId: nil,
	}

	admin := models.Role{
		Name:           "Администратор",
		Slug:           casbin.SubAdmin,
		Description:    &adminDesc,
		IsSystem:       true,
		OrganizationId: nil,
	}
	moder := models.Role{
		Name:           "Модератор",
		Slug:           casbin.SubModer,
		Description:    &moderatorDesc,
		IsSystem:       true,
		OrganizationId: nil,
	}
	teacher := models.Role{
		Name:           "Преподаватель",
		Slug:           casbin.SubTeacher,
		Description:    &teacherDesc,
		IsSystem:       true,
		OrganizationId: nil,
	}
	student := models.Role{
		Name:           "Студент",
		Slug:           casbin.SubStudent,
		Description:    &studentDesc,
		IsSystem:       true,
		OrganizationId: nil,
	}

	roleArray := []models.Role{
		owner,
		admin,
		moder,
		teacher,
		student,
	}

	for _, model := range roleArray {
		res := db.Create(&model)
		if res.Error != nil {
			panic(res.Error)
		}
	}

	return &roleArray

	//db.Transaction(func(tx *gorm.DB) error {
	//	for _, model := range roleArray {
	//		res := db.Create(&model)
	//		if res.Error != nil {
	//			return res.Error
	//		}
	//	}
	//	return nil
	//})
	//return &roleArray
}

func createOrganization(db *gorm.DB, owner models.User) models.Organization {
	organizationMembers := []models.User{owner}
	organization := models.Organization{
		Title:  "Example Org",
		Domain: "example.org",
		Email:  "example@organization.org",
		Active: true,
		User:   &organizationMembers,
	}
	db.Create(&organization)
	return organization
}

func createUsers(db *gorm.DB) *models.User {
	userRepo := repositories.NewUserRepository()
	hashPassword, err := helpers.CreatePasswordHash("admin")
	if err != nil {
		panic(err)
	}

	err, user := userRepo.CreateNewUser(&repositories.UserDataPayload{
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
	return user
}
