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

func AddEvent(name, description string, coach_id, gym_id sql.NullInt64, eventTime, eventDate string) error {

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var event_id int

	query := "INSERT INTO events (name, description, coach_id) VALUES ($1, $2, $3) RETURNING event_id"
	err = db.QueryRow(query, name, description, coach_id).Scan(&event_id)

	if err != nil {
		return err
	}

	query = "INSERT INTO current_events (gym_id, event_id, time, data) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, gym_id, event_id, eventTime, eventDate)

	if err != nil {
		return err
	}

	return nil
}
