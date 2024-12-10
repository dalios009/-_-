package main

import "fmt"

// PredictMatch provides a simple prediction between two teams based on team stats
func PredictMatch(homeTeamName, awayTeamName string) string {
	// Fetch team IDs based on names from a predefined map
	homeTeamID, ok := teamIDMap[homeTeamName]
	if !ok {
		return fmt.Sprintf("Team name %s not found", homeTeamName)
	}

	awayTeamID, ok := teamIDMap[awayTeamName]
	if !ok {
		return fmt.Sprintf("Team name %s not found", awayTeamName)
	}

	// Fetch team stats
	homeStats, err := FetchTeamStats(homeTeamID)
	if err != nil {
		return fmt.Sprintf("Error fetching stats for %s: %v", homeTeamName, err)
	}

	awayStats, err := FetchTeamStats(awayTeamID)
	if err != nil {
		return fmt.Sprintf("Error fetching stats for %s: %v", awayTeamName, err)
	}

	// Simple prediction based on team founding date
	if homeStats.Founded > awayStats.Founded {
		return fmt.Sprintf("%s is predicted to win against %s!", homeStats.Name, awayStats.Name)
	} else if awayStats.Founded > homeStats.Founded {
		return fmt.Sprintf("%s is predicted to win against %s!", awayStats.Name, homeStats.Name)
	} else {
		return fmt.Sprintf("It's a close match between %s and %s!", homeStats.Name, awayStats.Name)
	}
}

// teamIDMap is a sample mapping from team names to IDs
// Populate this with actual team IDs from your API or dataset
var teamIDMap = map[string]int{
	"Liverpool":       40, // Example ID; replace with actual ID for Liverpool
	"Manchester City": 50, // Example ID; replace with actual ID for Manchester City
	// Add more teams as needed
}
