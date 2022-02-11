package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// EnvConfig .
type EnvConfig struct {
	Debug string `envconfig:"DEBUG"`

	ServiceName string `envconfig:"SERVICE_NAME" required:"true"`
	ServerPort  int    `envconfig:"SERVER_PORT" default:"7723" required:"true"`

	MySQLDebug           bool   `envconfig:"MYSQL_DEBUG"`
	MySQLUser            string `envconfig:"MYSQL_USER"`
	MySQLPassword        string `envconfig:"MYSQL_PASSWORD"`
	MySQLHost            string `envconfig:"MYSQL_HOST"`
	MySQLPort            int    `envconfig:"MYSQL_PORT"`
	MySQLDatabase        string `envconfig:"MYSQL_DATABASE"`
	MySQLMaxOpenConn     int    `envconfig:"MYSQL_MAX_OPEN_CONNECTION"`
	MySQLMaxIdleConn     int    `envconfig:"MYSQL_MAX_IDLE_CONNECTION"`
	MySQLMaxLifetimeConn int    `envconfig:"MYSQL_MAX_LIFETIME_CONNECTION"`
}

// New .
func New() (*EnvConfig, error) {
	configuration := EnvConfig{}
	if err := readEnvFile(&configuration); err != nil {
		log.Fatal(err)
	}
	return &configuration, nil
}

func readEnvFile(object interface{}) error {
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", object); err != nil {
			return errors.Wrap(err, "failed to read from env variable")
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", object); err != nil {
		return errors.Wrap(err, "failed to read from env variable")
	}

	return nil
}
