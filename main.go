package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/getsentry/sentry-go"
	"github.com/jmoiron/sqlx"
	"github.com/millyleadley/roastadvisor/lib/log"
	"github.com/millyleadley/roastadvisor/lib/router"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	// Import the postgres driver.
	_ "github.com/lib/pq"
)

var (
	app        = kingpin.New("app", "Roast Advisor")
	configFile = app.Flag("config-file", "Which config file to use").Default("dev").Enum("dev", "prod")
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Create a context that can be cancelled.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Load the config file, which means we can access secrets throughout the app.
	loadConfig(*configFile)

	// Connect to the db.
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=postgres sslmode=disable", viper.GetString("db.user"), viper.GetString("db.password"))
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Error(errors.Wrap(err, "connecting to db"))
		return
	}
	defer db.Close()

	// Create the Sentry hub
	sentryDSN := viper.GetString("sentry.dsn")
	if sentryDSN != "" {
		opts := sentry.ClientOptions{
			Dsn:              string(sentryDSN),
			Environment:      *configFile,
			AttachStacktrace: true,
		}
		if err := sentry.Init(opts); err != nil {
			log.Error(errors.Wrap(err, "initialising Sentry"))
		}
		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
	} else {
		log.Warn(errors.New("no Sentry DSN found in config file"))
	}

	// Start the router - this should always be last because it blocks the main thread.
	router.Start(ctx, db)
}

func loadConfig(env string) {
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/hidden")

	if err := viper.ReadInConfig(); err != nil {
		log.Error(errors.Wrap(err, "reading config file"))
	}
}
