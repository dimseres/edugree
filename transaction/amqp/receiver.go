package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Receiver struct {
	connection *amqp.Connection
}

func NewAmqpReceiver(config *Config) *Receiver {
	conn, err := getConnection(config)
	if err != nil {
		failOnError(err, "Failed on creating new Receiver")
	}
	receiver := Receiver{connection: conn}
	return &receiver
}

type Message struct {
	EventName string
	Message   interface{}
}

func (receiver *Receiver) ListenQueue(config *QueueConfig, channel chan Message) {
	ch, err := receiver.connection.Channel()
	failOnError(err, "Failed to open a receiver channel")
	//defer ch.Close()

	queue, err := ch.QueueDeclare(
		config.Name,         // name
		config.Durable,      // durable
		config.UnusedDelete, // delete when unused
		config.Exclusive,    // exclusive
		config.NoWait,       // no-wait
		nil,
	)
	failOnError(err, "Failed to declare a receiver queue")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	go func() {
		for msg := range msgs {
			message := Message{
				EventName: msg.Type,
				Message:   msg.Body,
			}
			channel <- message
		}
	}()
}
