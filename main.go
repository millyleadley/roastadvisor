package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/spf13/viper"
)

var (
	app = kingpin.New("app", "Roast Advisor")
	env = app.Flag("env", "Which environment the app is being ran in").Default("dev").Enum("dev", "prod")
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	loadConfig(*env)

	test := viper.Get("test")
	fmt.Println(test)
}

func loadConfig(env string) {
	fmt.Println("Loading config file for", env)
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/hidden")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
