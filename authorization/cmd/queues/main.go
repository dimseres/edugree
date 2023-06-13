package main

import (
	"authorization/config"
	"authorization/internal/transport/amqp"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Start Queue Listener")
	login := config.GetConfig("RABBIT_LOGIN")
	password := config.GetConfig("RABBIT_PASSWORD")
	host := config.GetConfig("RABBIT_HOST")

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
