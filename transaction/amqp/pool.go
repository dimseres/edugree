package amqp

import (
	"fmt"
)

func ListenQueues(login string, password string, host string) {
	conf := Config{
		Login:    login,
		Password: password,
		Host:     host,
	}
	queue := NewAmqpReceiver(&conf)
	sagaQueue := QueueConfig{
		Name:         "saga_coordinator",
		Durable:      true,
		UnusedDelete: false,
		Exclusive:    false,
		NoWait:       false,
	}

	sagaQueueChannel := make(chan Message)
	queue.ListenQueue(&sagaQueue, sagaQueueChannel)

	fmt.Println("AmqpListener Started press ctrl + c to exit")
	for {
		message := <-sagaQueueChannel
		go HandleMessage(message)
	}
}
