package main

import (
	"authorization/internal/models"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db, _ := gorm.Open(postgres.Open(os.Getenv("DSN")))
	gormadapter.TurnOffAutoMigrate(db)
	_, err := gormadapter.NewAdapterByDBWithCustomTable(db, models.Permissions{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&models.Role{},
		&models.Permissions{},
		//&models.Membership{}, // m2m table created inside structs
		&models.User{},
		&models.Organization{},
	)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
