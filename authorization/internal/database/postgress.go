package database

import (
	"authorization/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB

func InitPgConnection() *gorm.DB {
	fmt.Println("Try to Database connect")
	logger := logger.New(log.New(config.GetLogger().Out, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
			ParameterizedQueries:      true,
		},
	)
	conn, err := gorm.Open(postgres.Open(config.GetConfig("DSN")), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		config.GetLogger().Error(err)
		panic(err)
	}
	db = conn
	return db
}

func GetConnection() *gorm.DB {
	return db
}
