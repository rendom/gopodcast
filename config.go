package main

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool

	DBFile string // Path to sqlite db TODO: support multiple db types

	JWTPublicKeyFile  string
	JWTPrivateKeyFile string
	JWTExpire         time.Duration
}

func GetConfig() *Config {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("/etc/gopodcast/")
	config.AddConfigPath("$HOME/.config/gopodcast")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Unable to load config %s\n", err)
	}

	var c Config
	err = config.Unmarshal(&c)

	if err != nil {
		log.Fatalf("Unable to unmarshal config: %s\n", err)
	}

	return &c
}
