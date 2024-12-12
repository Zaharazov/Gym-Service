package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
	"errors"
	"fmt"
)

func FetchAdmins() ([]models.Admin, error) {

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM admins"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []models.Admin
	for rows.Next() {
		var admin models.Admin
		if err := rows.Scan(&admin.ID, &admin.Login, &admin.Password, &admin.Name, &admin.Phone, &admin.GymID); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func FindAdmin(username string) (bool, string, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return false, "", err
	}
	defer db.Close()

	var password string

	query := `SELECT password FROM admins WHERE login = $1 LIMIT 1;
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

func CountAdmins() (int, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int

	query := `SELECT COUNT(admin_id) FROM admins`

	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

func AddAdmin(login, password, name string, phone string, gymID sql.NullInt64) error {

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

	// Запрос для вставки данных администратора
	query = `INSERT INTO admins (login, password, name, phone, gym_id) VALUES ($1, $2, $3, $4, $5)`

	// Выполнение запроса с параметрами
	_, err = db.Exec(query, login, password, name, phone, gymID)

	if err != nil {
		return err
	}

	return nil
}
