package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
	"errors"
	"fmt"
)

func FetchUsers() ([]models.User, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.Active, &user.Name, &user.Age, &user.Sex, &user.Phone, &user.PassID, &user.GymID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func AddUser(name, username, password string, age int, sex, phone string) error {

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	query := "SELECT EXISTS ( SELECT 1 FROM users WHERE login = $1 UNION SELECT 1 FROM admins WHERE login = $1 UNION SELECT 1 FROM coaches WHERE login = $1)"
	err = db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println(err)
		return errors.New("username already exists")
	}

	query = "INSERT INTO users (name, login, password, age, sex, phone) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = db.Exec(query, name, username, password, age, sex, phone)
	if err != nil {
		return err
	}

	return nil
}

func FindUser(username string) (bool, string, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false, "", err
	}
	defer db.Close()

	var password string

	query := `SELECT password FROM users WHERE login = $1 LIMIT 1;`

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

func CountUsers() (int, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int

	query := `SELECT COUNT(user_id) FROM users`

	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}
