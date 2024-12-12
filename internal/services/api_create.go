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
	coachIDStr := r.FormValue("coach_id")

	var nullableCoachID sql.NullInt64
	if coachIDStr != "" {
		parsedCoachID, err := strconv.Atoi(coachIDStr)
		if err == nil {
			nullableCoachID = sql.NullInt64{Int64: int64(parsedCoachID), Valid: true}
		}
	}

	gymIDStr := r.FormValue("gym_id")

	var nullableGymID sql.NullInt64
	if gymIDStr != "" {
		parsedGymID, err := strconv.Atoi(gymIDStr)
		if err == nil {
			nullableGymID = sql.NullInt64{Int64: int64(parsedGymID), Valid: true}
		}
	}

	time := r.FormValue("time")
	date := r.FormValue("date")

	err := database.AddEvent(name, description, nullableCoachID, nullableGymID, time, date)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/events", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/events", http.StatusFound)
}

func CreateCoach(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	sex := r.FormValue("sex")
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
	phone := r.FormValue("phone")
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
	err := database.AddAdmin(login, hashedPasswordStr, name, phone, nullableGymID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	// Перенаправление на страницу администрирования после успешного добавления
	http.Redirect(w, r, "/admin", http.StatusFound)
}
