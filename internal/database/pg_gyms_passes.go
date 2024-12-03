package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
)

func FetchGymsPasses() ([]models.Pass, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT gp_id, name, price, description FROM gym_passes"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gyms_passes []models.Pass
	for rows.Next() {
		var gym_pass models.Pass
		if err := rows.Scan(&gym_pass.ID, &gym_pass.Name, &gym_pass.Price, &gym_pass.Description); err != nil {
			return nil, err
		}
		gyms_passes = append(gyms_passes, gym_pass)
	}

	return gyms_passes, nil
}
