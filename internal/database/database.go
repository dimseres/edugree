package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func InitConnection() *gorm.DB {
	fmt.Println("Try to Database connect")
	conn, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
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
