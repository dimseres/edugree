package main

import (
	"authorization/config"
	casbin2 "authorization/internal/casbin"
	"authorization/internal/database"
	"authorization/internal/transport/rest"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

// start service
func main() {
	file := config.InitLogger()
	defer file.Close()
	connection := database.InitPgConnection()
	database.InitRedisConnection(&database.RedisConnectionConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	casbin2.InitCasbin(connection)

	fmt.Println(connection)
	rest.StartHttpServer("7001")
}
