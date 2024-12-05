package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
)

func FetchEquipment() ([]models.Equipment, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT equipment_id, name, description FROM fitness_equipment"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []models.Equipment
	for rows.Next() {
		var equipment models.Equipment
		if err := rows.Scan(&equipment.ID, &equipment.Name, &equipment.Description); err != nil {
			return nil, err
		}
		equipments = append(equipments, equipment)
	}

	return equipments, nil
}
