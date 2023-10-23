package log

import (
	"os"

	logorus "github.com/sirupsen/logrus"
)

var log *logorus.Logger

func init() {
	log = logorus.New()
	// log.SetFormatter(&logorus.JSONFormatter{})
	log.Out = os.Stdout
}

func GetLog() *logorus.Logger {
	return log
}
