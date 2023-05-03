package controllers

import (
	"transaction/amqp"
	"transaction/constants"
)

func Init() {
	amqp.RegisterEventHandler(constants.CREATE_ORGANIZATION_EVENT, func(c amqp.Message) error {
		return nil
	})
}
