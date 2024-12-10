package main

import (
	"log"
)

func main() {
	// Load the configuration (API keys, tokens, etc.)
	err := LoadConfig() // Load the keys from config.go
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Start the Telegram bot
	StartBot() // Start the bot from bot.go
}
