package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
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
		if err := rows.Scan(&admin.ID, &admin.Login, &admin.Password, &admin.Name, &admin.AccessLevel, &admin.Phone, &admin.GymID); err != nil {
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
