package amqp

import (
	"errors"
	"transaction/config"
)

type EventHandler func(message Message) error

var events = make(map[string]EventHandler)

func HandleMessage(message Message) error {
	handler, ok := events[message.EventName]
	if !ok {
		config.GetLogger().Error("Handler not found")
		return errors.New("handler not found")
	}

	return handler(message)
}

func RegisterEventHandler(eventName string, handler EventHandler) {
	events[eventName] = handler
}
