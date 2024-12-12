package database

import (
	"Gym-Service/internal/domain/models"
	"Gym-Service/internal/server/configs"
	"database/sql"
	"fmt"
)

func FetchGyms() ([]models.Gym, error) {
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
	query := "SELECT * FROM coaches WHERE gym_id = $1"

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
		if err := rows.Scan(&coach.ID, &coach.Login, &coach.Password, &coach.Name, &coach.Age, &coach.Sex, &coach.Description, &coach.GymID); err != nil {
			return nil, err
		}
		coaches = append(coaches, coach)
	}

	return coaches, nil
}

func FetchEquipmentsForGym(id string) ([]models.Equipment, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT b.equipment_id, b.name, b.description FROM availability as a INNER JOIN fitness_equipment as b ON a.equipment_id = b.equipment_id WHERE a.gym_id = $1"

	// Выполняем запрос
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
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

func FetchEventsForGym(id string) ([]models.Event, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT b.event_id, b.name, b.description, b.coach_id FROM current_events as a INNER JOIN events as b ON a.event_id = b.event_id WHERE a.gym_id = $1"

	// Выполняем запрос
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
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

func FetchUsersForGym(id string) ([]models.User, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT * FROM users WHERE gym_id = $1"

	// Выполняем запрос
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.Name, &user.Age, &user.Sex, &user.Phone, &user.PassID, &user.GymID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FetchAdminsForGym(id string) ([]models.Admin, error) {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// SQL-запрос для получения данных
	query := "SELECT * FROM admins WHERE gym_id = $1"

	// Выполняем запрос
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Читаем данные
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

func SetUserGymID(role, login string, gymID sql.NullInt64) error {

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var query string

	if role == "admin" {
		query = "UPDATE admins SET gym_id = $1 WHERE login = $2"
	} else if role == "user" {
		query = "UPDATE users SET gym_id = $1 WHERE login = $2"
	}

	_, err = db.Exec(query, gymID, login)
	if err != nil {
		return err
	}

	return nil
}
