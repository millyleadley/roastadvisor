package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/millyleadley/roastadvisor/lib/log"
	"github.com/spf13/viper"
)

var (
	app = kingpin.New("app", "Roast Advisor")
	env = app.Flag("env", "Which environment the app is being ran in").Default("dev").Enum("dev", "prod")
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Create a context that can be cancelled.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Load the config file, which means we can access secrets throughout the app.
	loadConfig(*env)
	test := viper.Get("test")
	log.Info("Loaded test value from config file", map[string]interface{}{"test": test})
}

func loadConfig(env string) {
	fmt.Println("Loading config file for", env)
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/hidden")

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file", err)
	}
}
