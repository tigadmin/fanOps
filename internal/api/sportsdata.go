package api

import (
	"encoding/json"
	"fmt"
	"github.com/your-username/fanOps/internal/database"
	"net/http"
)

// SportsData API endpoint for fetching teams
const SportsDataAPIURL = "https://api.sportsdata.io/v3/nba/scores/json/Teams?key=your_api_key"

// FetchTeams fetches the teams from the SportsData.io API and stores them in PostgreSQL
func FetchTeams() ([]database.Team, error) {
	resp, err := http.Get(SportsDataAPIURL)
	if err != nil {
		return nil, fmt.Errorf("could not fetch data: %v", err)
	}
	defer resp.Body.Close()

	var teams []database.Team
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		return nil, fmt.Errorf("could not parse response: %v", err)
	}

	// Store each team in the PostgreSQL database
	for _, team := range teams {
		if err := database.StoreTeam(team); err != nil {
			return nil, fmt.Errorf("could not store team in database: %v", err)
		}
	}

	return teams, nil
}
