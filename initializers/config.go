package initializers

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/viper"
)

func Setup() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if _, err := os.Stat("config.yml"); err != nil {
		buildInitialConfig()
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	rand.Seed(time.Now().UTC().UnixNano())
}

func buildInitialConfig() {

	file, err := os.Create("config.yml")
	if err != nil {
		log.Fatal("Couldn't create config file.")
	}
	file.Close()

	viper.SetDefault("dbUsername", "boardgames")
	viper.SetDefault("dbPassword", "QLpqd2p9nnC83eQjT53K5mLSqkJEY3Wf")
	viper.SetDefault("dbIP", "localhost")
	viper.SetDefault("dbPort", 5432)
	viper.SetDefault("dbName", "boardgames")
	viper.SetDefault("production", true)
	viper.SetDefault("apiPort", 8080)

	viper.WriteConfig()
}
