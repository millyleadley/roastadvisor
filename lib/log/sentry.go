package log

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

type SentryTag string

const (
	SentryTagEnvironment SentryTag = "environment"
)

// sendToSentry sends an error to Sentry.
func sendToSentry(err error, level sentry.Level, params ...map[string]interface{}) {
	// Whilst we're in dev, let's not do this.
	return

	if err != nil {
		hub := sentry.CurrentHub()

		// Build a sentry event.
		event := sentry.NewEvent()
		// Use the name of the error as the high-level name for the event in Sentry.
		event.Message = err.Error()
		event.Level = level

		// Add required metadata
		env := viper.GetString("env")
		hub.Scope().SetTag(string(SentryTagEnvironment), env)

		// Add any extra metadata
		for _, param := range params {
			for key, value := range param {
				event.Extra[key] = value
			}
		}

		hub.CaptureEvent(event)
	}
}
