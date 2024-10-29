package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/getsentry/sentry-go"
	"github.com/jmoiron/sqlx"
	reviewsServer "github.com/millyleadley/roastadvisor/api/gen/http/reviews/server"
	"github.com/millyleadley/roastadvisor/api/gen/reviews"
	apireviews "github.com/millyleadley/roastadvisor/app/reviews/api"
	"github.com/millyleadley/roastadvisor/lib/errors"
	"github.com/millyleadley/roastadvisor/lib/log"
	"github.com/spf13/viper"
	goahttp "goa.design/goa/v3/http"

	// Import the postgres driver, otherwise sqlx.Connect will fail.
	_ "github.com/lib/pq"
)

var (
	app        = kingpin.New("app", "Roast Advisor")
	configFile = app.Flag("config-file", "Which config file to use").Default("dev").Enum("dev", "prod")
)

func main() {
	log.Info("Starting app")

	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Create a context that can be cancelled.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Set up channel to handle graceful shutdown on interrupt signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Load the config file, which means we can access secrets throughout the app.
	err := loadConfig(ctx, *configFile)
	if err != nil {
		log.Error(errors.Wrap(ctx, err, "loading config"))
		return
	}
	log.Info("Config loaded")

	// Connect to the db.
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=postgres sslmode=disable", viper.GetString("db.user"), viper.GetString("db.password"))
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Error(errors.Wrap(ctx, err, "connecting to db"))
		return
	}
	defer db.Close()
	log.Info("Connected to db")

	// Create the Sentry hub
	sentryDSN := viper.GetString("sentry.dsn")
	if sentryDSN != "" {
		opts := sentry.ClientOptions{
			Dsn:              string(sentryDSN),
			Environment:      *configFile,
			AttachStacktrace: true,
		}
		if err := sentry.Init(opts); err != nil {
			log.Error(errors.Wrap(ctx, err, "initialising Sentry"))
		}
		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
	} else {
		log.Warn(errors.New(ctx, "no Sentry DSN found in config file"))
	}

	// Create Goa service and initialize endpoints.
	// TODO: How can we make this easier to add new services? What really is mux?
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	reviewsService := apireviews.NewReviews(db)
	reviewsEndpoints := reviews.NewEndpoints(reviewsService)
	reviewsSvr := reviewsServer.New(reviewsEndpoints, mux, dec, enc, nil, nil)
	reviewsSvr.Mount(mux)

	// TODO: Goa middleware can be added here.

	// Start the HTTP server.
	server := &http.Server{Addr: ":8000", Handler: mux}

	go func() {
		log.Info("Starting HTTP server on port 8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(errors.Wrap(ctx, err, "HTTP server error"))
		}
	}()

	// Wait for termination signal for graceful shutdown.
	<-stop
	log.Info("Shutting down server...")

	// Shutdown the server gracefully with a timeout.
	ctx, cancelTimeout := context.WithTimeout(ctx, 10*time.Second)
	defer cancelTimeout()
	if err := server.Shutdown(ctx); err != nil {
		log.Error(errors.Wrap(ctx, err, "Server forced to shutdown"))
	}

	log.Info("Server exited properly")
}

func loadConfig(ctx context.Context, env string) error {
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/hidden")

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(ctx, err, "reading config file")
	}

	return nil
}
