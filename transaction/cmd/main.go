package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"transaction/amqp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Start Queue Listener")
	login := os.Getenv("RABBIT_LOGIN")
	password := os.Getenv("RABBIT_PASSWORD")
	host := os.Getenv("RABBIT_HOST")

	amqp.ListenQueues(login, password, host)
}
