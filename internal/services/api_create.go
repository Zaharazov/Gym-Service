package services

import (
	"Gym-Service/internal/database"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	description := r.FormValue("description")
	id, _ := strconv.Atoi(r.FormValue("coach_id"))

	err := database.AddEvent(name, description, id)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/events", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/events", http.StatusFound)
}

func CreateCoach(w http.ResponseWriter, r *http.Request) {
	// Извлечение обязательных параметров
	login := r.FormValue("login")
	password := r.FormValue("password")
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	sex := r.FormValue("sex")

	// Извлечение необязательных параметров
	description := r.FormValue("description")
	gymIDStr := r.FormValue("gym_id")

	// Преобразование gym_id в sql.NullInt64
	var nullableGymID sql.NullInt64
	if gymIDStr != "" {
		parsedGymID, err := strconv.Atoi(gymIDStr)
		if err == nil {
			nullableGymID = sql.NullInt64{Int64: int64(parsedGymID), Valid: true}
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)

	// Передача данных во внутреннюю функцию
	err := database.AddCoach(login, hashedPasswordStr, name, age, sex, description, nullableGymID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	// Перенаправление после успешного выполнения
	http.Redirect(w, r, "/admin", http.StatusFound)
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	// Сбор обязательных данных из формы
	login := r.FormValue("login")
	password := r.FormValue("password")
	name := r.FormValue("name")
	accessLevel, _ := strconv.Atoi(r.FormValue("access_level"))
	phone := r.FormValue("phone")

	// Сбор необязательных данных
	gymIDStr := r.FormValue("gym_id")

	var nullableGymID sql.NullInt64
	if gymIDStr != "" {
		parsedGymID, err := strconv.Atoi(gymIDStr)
		if err == nil {
			nullableGymID = sql.NullInt64{Int64: int64(parsedGymID), Valid: true}
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)

	// Вызов функции добавления администратора в базу данных
	err := database.AddAdmin(login, hashedPasswordStr, name, accessLevel, phone, nullableGymID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	// Перенаправление на страницу администрирования после успешного добавления
	http.Redirect(w, r, "/admin", http.StatusFound)
}
