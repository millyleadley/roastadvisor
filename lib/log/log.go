package log

// A thin wrapper around logrus.

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func Info(msg string, params ...map[string]interface{}) {
	loggerWithMetadata(params).Info(msg)
}

func Warn(err error, params ...map[string]interface{}) {
	sendToSentry(err, sentry.LevelWarning, params...)
	loggerWithMetadata(params).WithError(err).Warn()
}

func Error(err error, params ...map[string]interface{}) {
	sendToSentry(err, sentry.LevelError, params...)
	loggerWithMetadata(params).WithError(err).Error()
}

// loggerWithMetadata returns a logrus logger with the params as fields.
func loggerWithMetadata(params []map[string]interface{}) *logrus.Entry {
	fields := logrus.Fields{}
	for _, param := range params {
		for key, value := range param {
			fields[key] = value
		}
	}
	return logger.WithFields(fields)
}
