package logger

import "github.com/sirupsen/logrus"

var (
	Log = logrus.WithFields(logrus.Fields{
		"lib": "grule-rule-engine",
	})
)

func init() {
	Log.Level = logrus.InfoLevel
}

// SetLogLevel will set the logrus log level
func SetLogLevel(lvl logrus.Level) {
	Log.Level = lvl
}
