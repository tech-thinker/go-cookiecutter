package logger

import "github.com/sirupsen/logrus"

// Log is public object for log handling
var Log *logrus.Logger

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	Log = logrus.New()
}
