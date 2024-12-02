package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
	"fmt"
)

type PageData struct {
	Gyms []models.Gym
	Role string
}

func FetchGymsFromDB() ([]models.Gym, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT gym_id, name, address, description FROM gyms"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gyms []models.Gym
	for rows.Next() {
		var gym models.Gym
		if err := rows.Scan(&gym.ID, &gym.Name, &gym.Address, &gym.Description); err != nil {
			return nil, err
		}
		gyms = append(gyms, gym)
	}

	return gyms, nil
}

func FetchGymByID(id string) (models.Gym, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return models.Gym{}, err
	}
	defer db.Close()

	// SQL-запрос для получения данных о зале по ID
	query := "SELECT gym_id, name, address, description FROM gyms WHERE gym_id = $1"

	// Используем QueryRow, так как ожидаем только одну строку результата
	row := db.QueryRow(query, id)

	// Читаем данные из результата
	var gym models.Gym
	err = row.Scan(&gym.ID, &gym.Name, &gym.Address, &gym.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если зал с указанным ID не найден
			return models.Gym{}, fmt.Errorf("gym with ID %s not found", id)
		}
		// Другая ошибка
		return models.Gym{}, err
	}

	// Возвращаем найденный зал
	return gym, nil
}

func FetchTrainersForGym(id string) ([]models.Coach, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT name, description FROM coaches WHERE gym_id = $1"

	// Выполняем запрос
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
	var coaches []models.Coach
	for rows.Next() {
		var coach models.Coach
		if err := rows.Scan(&coach.Name, &coach.Description); err != nil {
			return nil, err
		}
		coaches = append(coaches, coach)
	}

	return coaches, nil
}
