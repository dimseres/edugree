package amqp

import (
	"transaction/config"
	"transaction/constants"
)

type Config struct {
	Login    string
	Password string
	Host     string
}

type QueueConfig struct {
	Name         string
	Durable      bool
	UnusedDelete bool
	Exclusive    bool
	NoWait       bool
}

var EventQueueMap = map[string][]string{
	constants.CREATE_ORGANIZATION_EVENT: {"course_queue"},
	constants.CREATE_USERS:              {"course_queue"},
}

func failOnFireError(event string, err error, message *QueueMessage) {
	if err != nil {
		config.GetLogger().Error(err, message.Type, message.Message)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		config.GetLogger().Error(err, msg)
		//log.Panicf("%s: %s", msg, err)
	}
}
