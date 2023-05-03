package amqp

import (
	"errors"
	"transaction/config"
)

type EventHandler func(c Message) error

var events = make(map[string]EventHandler)

func HandleMessage(message Message) error {
	handler, ok := events[message.EventName]
	if !ok {
		config.GetLogger().Error("Handler not found")
		return errors.New("Handler not found")
	}

	return handler(message)
}

func RegisterEventHandler(eventName string, handler EventHandler) {
	events[eventName] = handler
}
