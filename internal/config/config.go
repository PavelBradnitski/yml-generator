package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config хранит всю конфигурацию приложения.
type Config struct {
	BaseURL        string
	APIURL         string
	APIToken       string
	OutputFilename string
}

// Load загружает конфигурацию из .env файла и переменных окружения.
func Load() (*Config, error) {

	godotenv.Load()

	cfg := &Config{
		BaseURL:        os.Getenv("BASE_URL"),
		APIURL:         os.Getenv("API_URL"),
		APIToken:       os.Getenv("API_TOKEN"),
		OutputFilename: os.Getenv("OUTPUT_FILENAME"),
	}

	if cfg.APIURL == "" {
		return nil, errors.New("переменная окружения API_URL не установлена")
	}
	if cfg.APIToken == "" {
		return nil, errors.New("переменная окружения API_TOKEN не установлена")
	}
	if cfg.OutputFilename == "" {
		cfg.OutputFilename = "shop.yml"
	}

	return cfg, nil
}
