package amqp

import "log"

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
