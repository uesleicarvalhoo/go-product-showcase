package logger

import (
	"github.com/sirupsen/logrus"
)

func InitLogger(debug bool) {
	level := logrus.InfoLevel
	if debug {
		level = logrus.DebugLevel
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(level)
}
