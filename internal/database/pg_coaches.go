package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
	"errors"
	"fmt"
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

func AddCoach(login, password, name string, age int, sex string, description string, gymID sql.NullInt64) error {

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	query := "SELECT EXISTS ( SELECT 1 FROM users WHERE login = $1 UNION SELECT 1 FROM admins WHERE login = $1 UNION SELECT 1 FROM coaches WHERE login = $1)"
	err = db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println(err)
		return errors.New("login already exists")
	}

	// Запрос для вставки данных тренера
	query = `INSERT INTO coaches (login, password, name, age, sex, description, gym_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Выполнение запроса с параметрами
	_, err = db.Exec(query, login, password, name, age, sex, description, gymID)

	if err != nil {
		return err
	}

	return nil
}
