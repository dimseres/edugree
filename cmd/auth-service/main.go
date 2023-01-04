package main

import (
	"edugree_auth/internal"
	"edugree_auth/internal/database"
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
	internal.StartHttpServer("7001")
}
