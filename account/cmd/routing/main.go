package main

import (
	"account/database"
	rest "account/internal/transport/http"
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
	rest.StartHttpServer("7000")
}
