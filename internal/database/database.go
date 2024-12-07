package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// InitDB initializes the connection to the PostgreSQL database
func InitDB() (*sql.DB, error) {
	// Update this connection string with your actual database credentials
	connStr := "postgres://vultradmin:password@hostname:port/dbname"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	// Assign the DB connection to the global variable
	DB = db
	return DB, nil
}

// StoreTeam stores a team in the PostgreSQL database
func StoreTeam(team Team) error {
	query := `INSERT INTO teams (id, name, city, abbreviation, conference, division) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := DB.Exec(query, team.ID, team.Name, team.City, team.Abbreviation, team.Conference, team.Division)
	return err
}

type Team struct {
	ID           int
	Name         string
	City         string
	Abbreviation string
	Conference   string
	Division     string
}
