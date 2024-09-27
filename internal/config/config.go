//go:build !test

package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

// Config Common config struct
type Config struct {
	Service  Service
	Auth     Auth
	User     User
	Metrics  Metrics
	Platform Platform
}

// Service struct for storage this server config variables
type Service struct {
	Port string `env:"GATEWAY_SERVICE_PORT"`
}

// Auth struct for storage auth-service config variables
type Auth struct {
	Host string `env:"AUTH_SERVICE_HOST"`
	Port string `env:"AUTH_SERVICE_PORT"`
}

type User struct {
	Host string `env:"USER_SERVICE_HOST"`
	Port string `env:"USER_SERVICE_PORT"`
}

type Metrics struct {
	Host string `env:"GRAFANA_HOST"`
	Port int    `env:"GRAFANA_PORT"`
}

type Platform struct {
	Secret string `env:"SECRET_KEY"`
	Env    string `env:"ENV"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("Can not read env variables: %s", err)
	}
	return cfg
}
