package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func StartBot() {
	// Set up the bot with the Telegram API token
	pref := tele.Settings{
		Token:  config.TelegramBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	// Handle /start command
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Welcome to the Football Predictor Bot! Use /player [name] to search for a player. Use /predict [team1] vs [team2] to get a match prediction.")
	})

	// Handle /player command for player search
	bot.Handle("/player", func(c tele.Context) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Send("Please provide a player name to search for.")
		}

		query := strings.Join(args, " ")
		players, err := SearchPlayers(query)
		if err != nil {
			return c.Send(fmt.Sprintf("Error fetching player data: %v", err))
		}

		if len(players) == 0 {
			return c.Send("No players found.")
		}

		// Build response with player information
		var response string
		for _, player := range players {
			response += fmt.Sprintf("Name: %s, Team: %s\n", player.Name, player.TeamName)
		}
		return c.Send(response)
	})

	// Handle /predict command for match prediction
	bot.Handle("/predict", func(c tele.Context) error {
		args := c.Args()

		// We expect the format "team1 vs team2"
		if len(args) < 3 {
			return c.Send("Please use the format: /predict [team1] vs [team2].")
		}

		teams := strings.Join(args, " ")
		splitTeams := strings.Split(teams, " vs ")
		if len(splitTeams) != 2 {
			return c.Send("Please use the format: /predict [team1] vs [team2].")
		}

		homeTeam := strings.TrimSpace(splitTeams[0])
		awayTeam := strings.TrimSpace(splitTeams[1])

		// Call PredictMatch from analysis.go
		prediction := PredictMatch(homeTeam, awayTeam)
		return c.Send(fmt.Sprintf("Prediction: %s", prediction))
	})

	// Start the bot
	bot.Start()
}
