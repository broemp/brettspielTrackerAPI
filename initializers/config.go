package initializers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func Setup() {

	viper.SetDefault("dbUsername", "boardgames")
	viper.SetDefault("dbPassword", "QLpqd2p9nnC83eQjT53K5mLSqkJEY3Wf")
	viper.SetDefault("DBIP", "localhost")
	viper.SetDefault("dbPort", 5432)
	viper.SetDefault("dbName", "boardgames")
	viper.SetDefault("production", true)
	viper.SetDefault("apiPort", 8080)

	if _, err := os.Stat("config.yml"); !errors.Is(err, os.ErrNotExist) {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Config Error: %w", err))
		}
	}

	// Override Default Settings with env variables
	for _, env := range viper.AllKeys() {
		if val, present := os.LookupEnv(strings.ToUpper(env)); present {
			viper.Set(env, val)
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
}
