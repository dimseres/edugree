package amqp

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Client struct {
	connection *amqp.Connection
}

type QueueMessage struct {
	Exchange    string
	Type        string
	ContentType string
	Message     string
}

func (client *Client) SendMessage(config *QueueConfig, message *QueueMessage) {
	channel, err := client.connection.Channel()
	failOnError(err, "Failed to open chanel")
	queue, err := channel.QueueDeclare(
		config.Name,
		config.Durable,
		config.UnusedDelete,
		config.Exclusive,
		config.NoWait,
		nil,
	)
	failOnError(err, "Failed to declare Queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx, message.Exchange, queue.Name, false, false, amqp.Publishing{
		ContentType: message.ContentType,
		Body:        []byte(message.Message),
	})

	failOnError(err, "Failed to Publish a message")
}

func NewAmqpClient(config *Config) (*Client, error) {
	conn, err := getConnection(config)
	if err != nil {
		failOnError(err, "Failed on creating new Client")
	}
	client := Client{connection: conn}
	return &client, err
}

func getConnection(config *Config) (*amqp.Connection, error) {
	dsn := "amqp://" + config.Login + ":" + config.Password + "@" + config.Host
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return conn, err
}
