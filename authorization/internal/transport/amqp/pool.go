package amqp

import "fmt"

func ListenQueues(login string, password string, host string) {
	conf := Config{
		Login:    login,
		Password: password,
		Host:     host,
	}
	queue := NewAmqpReceiver(&conf)
	authQueue := QueueConfig{
		Name:         "auth_queue",
		Durable:      false,
		UnusedDelete: false,
		Exclusive:    false,
		NoWait:       false,
	}

	authQueueChannel, customQueueChannel := make(chan Message), make(chan Message)
	queue.ListenQueue(&authQueue, authQueueChannel)

	fmt.Println("AmqpListener Started press ctrl + c to exit")
	for {
		select {
		case authMessage := <-authQueueChannel:
			fmt.Println("MESSAGE FROM AUTH QUEUE \n", authMessage.Message)
		case customMessage := <-customQueueChannel:
			fmt.Println(customMessage.Message)
		}
	}
}
