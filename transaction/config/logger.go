package config

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger = log.New()

func InitLogger() *os.File {
	f, err := os.OpenFile("logs/log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(io.MultiWriter(f, os.Stdout))
	logger.SetLevel(log.DebugLevel)
	return f
}

func GetLogger() *log.Logger {
	return logger
}
