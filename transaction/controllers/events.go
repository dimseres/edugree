package controllers

import (
	"transaction/amqp"
	"transaction/constants"
)

func RegisterEvents() {
	amqp.RegisterEventHandler(constants.CREATE_ORGANIZATION_EVENT, func(message amqp.Message) error {
		return nil
	})
}
