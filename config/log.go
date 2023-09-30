package config

import "github.com/sirupsen/logrus"

func NewLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
