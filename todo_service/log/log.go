package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.Out = os.Stdout
}

func GetLog() *logrus.Logger {
	return log
}
