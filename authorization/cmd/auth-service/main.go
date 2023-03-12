package main

import (
	casbin2 "authorization/internal/casbin"
	"authorization/internal/database"
	"authorization/internal/transport/rest"
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	connection := database.InitConnection()
	casbin2.InitCasbin(connection)

	fmt.Println(connection)
	rest.StartHttpServer("7001")
}
