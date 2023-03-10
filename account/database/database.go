package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitConnection() *gorm.DB {
	fmt.Println("Try to Database connect")
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
			ParameterizedQueries:      true,
		},
	)
	conn, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}
	db = conn
	fmt.Println("Connected to DB")
	return db
}

func GetConnection() *gorm.DB {
	return db
}
