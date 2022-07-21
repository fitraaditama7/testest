package config

import (
	"log"

	"test/pkg/database"
	"test/pkg/env"
	"test/pkg/server"
)

const prefix = "PROJECT"

type Config struct {
	ServerConfig server.Options    `envconfig:"SERVER"`
	DBConfig     database.DBConfig `envconfig:"DATABASE"`
}

func LoadConfig() *Config {
	var config Config
	err := env.Load(prefix, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration. err: " + err.Error())
	}
	return &config
}
