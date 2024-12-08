package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
)

func FetchCoaches() ([]models.Coach, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM coaches"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coaches []models.Coach
	for rows.Next() {
		var coach models.Coach
		if err := rows.Scan(&coach.ID, &coach.Login, &coach.Password, &coach.Name, &coach.Age, &coach.Sex, &coach.Description, &coach.GymID); err != nil {
			return nil, err
		}
		coaches = append(coaches, coach)
	}

	return coaches, nil
}

func FindCoach(username string) (bool, string, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false, "", err
	}
	defer db.Close()

	var password string

	query := `SELECT password FROM coaches WHERE login = $1 LIMIT 1;
    `

	// Выполнение SQL-запроса.
	err = db.QueryRow(query, username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil // Пользователь не найден
		}
		return false, "", err // Ошибка базы данных
	}

	// Если строка найдена
	exists := true
	return exists, password, nil
}

func CountCoaches() (int, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int

	query := `SELECT COUNT(coach_id) FROM coaches`

	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}
