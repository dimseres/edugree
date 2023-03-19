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
	casbin.InitCasbin(connection)
	casbin.DefineInitialPolicies(casbin.TechincalDomain)

	connection.Transaction(func(tx *gorm.DB) error {
		roles := createRoles(connection)
		user := createUsers(connection)
		services := createServices(connection)
		organization := createOrganization(connection, *user)
		organization = *bindOrganizationServices(connection, &organization, services)
		_ = addMember(connection, *user, organization, roles[0])

		casbin.DefineInitialPolicies(organization.Domain)

		return nil
	})

}

func createRoles(db *gorm.DB) []*models.Role {
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

	roleArray := []*models.Role{
		&owner,
		&admin,
		&moder,
		&teacher,
		&student,
	}

	for _, model := range roleArray {
		res := db.Create(&model)
		if res.Error != nil {
			panic(res.Error)
		}
	}

	return roleArray

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
	//organizationMembers := []models.User{owner}
	organization := models.Organization{
		Title:  "Example Org",
		Domain: "example.org",
		Email:  "example@organization.org",
		Active: true,
	}
	res := db.Create(&organization)
	if res.Error != nil {
		panic(res.Error)
	}
	return organization
}

func addMember(db *gorm.DB, user models.User, organization models.Organization, role *models.Role) models.Membership {
	members := models.Membership{
		UserId:         user.Id,
		OrganizationId: organization.Id,
		RoleId:         role.Id,
	}
	res := db.Create(&members)
	if res.Error != nil {
		panic(res.Error)
	}
	return members
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

func createServices(db *gorm.DB) []*models.Service {
	courseDescr := "В данном сервисе присутсвует возможность редактировать и создавать курсы, назначать ответственных"
	courseService := models.Service{
		Title:       "Каталог курсов",
		Slug:        "courses",
		Description: &courseDescr,
		User:        nil,
	}

	messengerDescr := "Сервис отпарвки сообщений между членами организации"
	messengerService := models.Service{
		Title:       "Мессенджер",
		Slug:        "messenger",
		Description: &messengerDescr,
		User:        nil,
	}

	services := []*models.Service{
		&courseService,
		&messengerService,
	}

	for _, model := range services {
		res := db.Create(&model)
		if res.Error != nil {
			panic(res.Error)
		}
	}

	return services
}

func bindOrganizationServices(db *gorm.DB, organization *models.Organization, services []*models.Service) *models.Organization {
	updatedOrganization := organization
	srvc := []models.Service{}
	for _, service := range services {
		srvc = append(srvc, *service)
	}
	updatedOrganization.Services = &srvc
	res := db.Save(&organization)
	if res.Error != nil {
		panic(res.Error)
	}
	return updatedOrganization
}
