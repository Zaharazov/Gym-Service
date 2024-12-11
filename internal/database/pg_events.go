package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
)

func FetchEvents() ([]models.Event, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM events"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.CoachID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func AddEvent(name, description string, id int) error {

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO events (name, description, coach_id) VALUES ($1, $2, $3)"
	if id == 0 {
		_, err = db.Exec(query, name, description, nil)
	} else {
		_, err = db.Exec(query, name, description, id)
	}

	if err != nil {
		return err
	}

	return nil
}
