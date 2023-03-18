package main

import (
	"authorization/internal/models"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db, _ := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Silent,
		}),
	})
	gormadapter.TurnOffAutoMigrate(db)
	_, err := gormadapter.NewAdapterByDBWithCustomTable(db, models.Permissions{})
	if err != nil {
		panic(err)
	}
	//err = db.Debug().AutoMigrate(&models.User{})
	err = db.Debug().AutoMigrate(
		&models.Membership{},
		&models.User{},
		&models.Organization{},
		&models.Role{},
		&models.Permissions{},
		&models.Service{},
	)

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
