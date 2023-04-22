package main

import (
	"authorization/internal/transport/amqp"
	"fmt"
	"github.com/joho/godotenv"
	"os"
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

	//client, err := amqp.NewAmqpClient(&amqp.Config{
	//	Login:    login,
	//	Password: password,
	//	Host:     host,
	//})
	//if err != nil {
	//	panic(err)
	//}

	//client.SendMessage(&amqp.QueueConfig{
	//	Name:    "course_queue",
	//	Durable: true,
	//}, &amqp.QueueMessage{
	//	Type:        "organization.create",
	//	ContentType: "application/json",
	//	Message:     "message",
	//})

	amqp.ListenQueues(login, password, host)
}
