package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBUser     string `env:"DB_USER" required:"true"`
	DBPassword string `env:"DB_PASSWORD" required:"true"`
	DBHost     string `env:"DB_HOST" required:"true"`
	DBPort     string `env:"DB_PORT" required:"true"`
	DBName     string `env:"DB_NAME" required:"true"`

	GRPCPort string `env:"COMMENT_SERVICE_PORT" required:"true"`
}

func LoadConfig() *Config {
	path := "./.env"
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
