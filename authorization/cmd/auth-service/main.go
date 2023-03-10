package main

import (
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
	fmt.Println(connection)
	rest.StartHttpServer("7001")
}
