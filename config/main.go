package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type configType struct {
	Env             string
	IsEnvProduction bool
	Host            string
	Port            string
}

var (
	envFileLoaded bool
	config        *configType
)

func Get() configType {
	if config != nil {
		return *config
	}

	if !envFileLoaded {
		envFile := os.Getenv("ENV_FILE")
		if envFile == "" {
			envFile = ".env"
		}
		if _, err := os.Stat(envFile); !os.IsNotExist(err) {
			err := godotenv.Load(envFile)
			if err != nil {
				log.Fatalf("Error loading .env file: %s", err)
			}
		}

		envFileLoaded = true
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
		if strings.HasSuffix(os.Args[0], ".test") || strings.Contains(os.Args[0], "/_test/") {
			env = "testing"
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("LISTEN_ADDR")
		if port == "" {
			port = "4040"
		}
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = ":"
	}

	config = &configType{
		Env:             env,
		IsEnvProduction: env == "production",
		Host:            host,
		Port:            port,
	}

	return *config
}
