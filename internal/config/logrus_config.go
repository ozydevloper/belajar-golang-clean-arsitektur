package config

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(c ConfigLogger) *logrus.Logger {

	log := logrus.New()

	log.SetLevel(logrus.Level(c.Level))
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	return log
}
