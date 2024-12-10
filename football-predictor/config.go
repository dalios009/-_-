package main

import (
	"fmt"
	"os"
)

type Config struct {
	TelegramBotToken string
	RapidAPIKey      string
	FootballAPIHost  string
}

var config Config

// LoadConfig loads API keys from environment variables or hardcodes them for testing
func LoadConfig() error {
	config.TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	config.RapidAPIKey = os.Getenv("RAPIDAPI_KEY")
	config.FootballAPIHost = "free-api-live-football-data.p.rapidapi.com"

	if config.TelegramBotToken == "" {
		config.TelegramBotToken = "your_telegram_token"
	}
	if config.RapidAPIKey == "" {
		config.RapidAPIKey = "98b573a717mshf741403153ee97ep10f89djsn0fdef72d989b"
	}

	if config.TelegramBotToken == "" || config.RapidAPIKey == "" || config.FootballAPIHost == "" {
		return fmt.Errorf("missing required API keys or tokens")
	}
	return nil
}
