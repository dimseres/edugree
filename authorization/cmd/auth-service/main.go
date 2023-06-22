package main

import (
	"authorization/config"
	casbin2 "authorization/internal/casbin"
	"authorization/internal/database"
	"authorization/internal/transport/rest"
	"fmt"
)

func init() {
	config.InitConfigs()
}

// start service
func main() {
	file := config.InitLogger()
	defer file.Close()
	connection := database.InitPgConnection()
	database.InitRedisConnection(&database.RedisConnectionConfig{
		Host:     config.GetConfig("REDIS_HOST"),
		Password: config.GetConfig("REDIS_PASSWORD"),
	})
	casbin2.InitCasbin(connection)

	fmt.Println(connection)
	rest.StartHttpServer("7001")
}
