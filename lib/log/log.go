package log

// A thin wrapper around logrus.

import (
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
	logger.WithFields(paramsToFields(params)).Info(msg)
}

func Warn(msg string, err error, params ...map[string]interface{}) {
	logger.WithFields(paramsToFields(params)).WithError(err).Warn(msg)
}

func Error(msg string, err error, params ...map[string]interface{}) {
	logger.WithFields(paramsToFields(params)).WithError(err).Error(msg)
}

func paramsToFields(params []map[string]interface{}) logrus.Fields {
	fields := logrus.Fields{}
	for _, param := range params {
		for key, value := range param {
			fields[key] = value
		}
	}
	return fields
}
