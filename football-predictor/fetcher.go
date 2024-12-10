package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struct for team stats
type TeamStats struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Founded int    `json:"founded"`
}

// Struct for player data
type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsCoach  bool   `json:"isCoach"`
	TeamID   int    `json:"teamId"`
	TeamName string `json:"teamName"`
}

// FetchTeamStats fetches team statistics by team ID
func FetchTeamStats(teamID int) (*TeamStats, error) {
	url := fmt.Sprintf("https://free-api-live-football-data.p.rapidapi.com/teams/%d", teamID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("x-rapidapi-host", config.FootballAPIHost)
	req.Header.Add("x-rapidapi-key", config.RapidAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %s: %s", resp.Status, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var team TeamStats
	if err := json.Unmarshal(body, &team); err != nil {
		return nil, fmt.Errorf("error parsing response body: %v", err)
	}

	return &team, nil
}

// SearchPlayers searches for football players by name
func SearchPlayers(query string) ([]Player, error) {
	url := fmt.Sprintf("https://%s/football-players-search?search=%s", config.FootballAPIHost, query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("x-rapidapi-host", config.FootballAPIHost)
	req.Header.Add("x-rapidapi-key", config.RapidAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %s: %s", resp.Status, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Adjust the struct to match the actual API response structure
	var result struct {
		Status   string `json:"status"`
		Response struct {
			Suggestions []Player `json:"suggestions"`
		} `json:"response"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %v", err)
	}

	// Check if the status indicates success
	if result.Status != "success" {
		return nil, fmt.Errorf("API response status is not success: %s", result.Status)
	}

	return result.Response.Suggestions, nil
}
