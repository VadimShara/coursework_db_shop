package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env			string			`env:"ENV" env-default:"local"`
	LogLevel	string			`env:"LOG_LEVEL" env-default:"info"`

	Port		int				`env:"PORT" env-default:"true"`
	IdleTimeout	time.Duration	`env:"IDLE_TIMEOUT" env-default:"10s"`
	ReqTimeout	time.Duration	`env:"REQUEST_TIMEOUT" env-default:"5s"`

	PGConfig	PostgresConfig
}

type PostgresConfig struct {
	User		string			`env:"POSTGRES_USER" env-required:"true"`
	Password	string			`env:"POSTGRES_PASSWORD" env-required:"true"`
	Host		string			`env:"POSTGRES_HOST" env-required:"true"`
	Port		string			`env:"POSTGRES_PORT" env-required:"true"`
	DBName		string			`env:"POSTGRES_DBNAME" env-required:"true"`		
	SSLMode		string			`env:"POSTGRES_SSLMODE" env-required:"disable"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Load() Config {
	var config Config

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatal("couldn't bind settings to config")
	}

	return config
}