package services

import (
	"authorization/internal/constants"
	"authorization/internal/models"
	"authorization/internal/transport/amqp"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

type IEventClient interface {
	SendMessage(config *amqp.QueueConfig, message *amqp.QueueMessage)
}

type EventService struct {
	client IEventClient
}

type EventDTO struct {
	Type    string
	Payload string
}

func NewEventService() EventService {
	client, err := amqp.NewAmqpClient(&amqp.Config{
		Login:    os.Getenv("RABBIT_LOGIN"),
		Password: os.Getenv("RABBIT_PASSWORD"),
		Host:     os.Getenv("RABBIT_HOST"),
	})

	if err != nil {
		log.Printf("ERROR ON CREATING EVENT: %s", err.Error())
	}

	return EventService{client: client}
}

func (self *EventService) FireEvent(eventType string, payload interface{}) bool {
	var data string
	var err error
	switch eventType {
	case constants.CREATE_ORGANIZATION_EVENT:
		data, err = createOrganization(payload)
	case constants.CREATE_USER:
		data, err = createUserEvent(payload)

	}

	if err != nil || data == "" {
		return false
	}

	for _, queueName := range amqp.EventQueueMap[eventType] {
		self.client.SendMessage(&amqp.QueueConfig{
			Name:         queueName,
			Durable:      true,
			UnusedDelete: false,
			Exclusive:    false,
			NoWait:       false,
		}, &amqp.QueueMessage{
			Exchange:    "",
			Type:        eventType,
			ContentType: "application/json",
			Message:     data,
		})
	}

	return false
}

func toJsonString(data []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func createOrganization(payload interface{}) (string, error) {
	model := payload.(models.Organization)
	data, err := json.Marshal(model)
	return toJsonString(data, err)
}

func createUserEvent(payload interface{}) (string, error) {
	model := payload.(models.User)
	data, err := json.Marshal(model)
	return toJsonString(data, err)
}

func FireEvent(eventName string, payload interface{}) {
	eventService := NewEventService()
	eventService.FireEvent(eventName, payload)
}
