// package config is responsible to handle the configuration processes
package config

import (
	"log"

	"github.com/spf13/viper"
)

// Struct to hold the config values
type Config struct {
	Port      string `mapstructure:"PORT"`
	Ns1ApiKey string `mapstructure:"API_SECRET"`
	Ns1Host   string `mapstructure:"API_URL"`
	Timeout   int    `mapstructure:"TIMEOUT_SECONDS"`
}

// Initialize private config var
var config Config

// GetConfig checks if config was set, if so, returns config, otherwise use viper to set its values
func GetConfig() *Config {
	if config != (Config{}) {
		return &config
	}

	viper.SetConfigFile(".env")
	e := viper.ReadInConfig()
	if e != nil {
		log.Panic("Viper: Error while reading config")
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	if (Config{}) == config {
		panic("No .env file provided")
	}
	return &config
}
