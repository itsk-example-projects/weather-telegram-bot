package config

import (
	"log"
	"os"
	"strings"
)

const (
	TelegramBotToken = "TELEGRAM_BOT_TOKEN"
)

type Config struct {
	TelegramBotToken string
	WeatherApiKey    string
}

func Load() (*Config, error) {
	var missing []string

	botToken := os.Getenv(TelegramBotToken)
	if botToken == "" {
		missing = append(missing, TelegramBotToken)
	}

	if len(missing) > 0 {
		log.Fatalf("missing required environment variables: %s", strings.Join(missing, ", "))
		return nil, nil // Otherwise IDE swears
	} else {
		return &Config{TelegramBotToken: botToken}, nil
	}
}
