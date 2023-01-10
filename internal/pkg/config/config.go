package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	Ns1ApiKey string `mapstructure:"API_SECRET"`
	Ns1Host   string `mapstructure:"API_URL"`
	Timeout   int    `mapstructure:"TIMEOUT_SECONDS"`
}

var SDKConfig Config

func GetConfig() *Config {
	if SDKConfig != (Config{}) {
		return &SDKConfig
	}

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	err := viper.Unmarshal(&SDKConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	if (Config{}) == SDKConfig {
		panic("No .env file provided")
	}
	return &SDKConfig
}
