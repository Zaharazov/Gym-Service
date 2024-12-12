package services

import (
	"Gym-Service/internal/database"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func SetPass(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.ServeFile(w, r, "./frontend/pages/not_auth_page.html")
		return
	}

	username := session.Values["username"].(string)

	passIDStr := r.FormValue("pass_id")

	var nullablePassID sql.NullInt64
	if passIDStr != "" {
		parsedPassID, err := strconv.Atoi(passIDStr)
		if err == nil {
			nullablePassID = sql.NullInt64{Int64: int64(parsedPassID), Valid: true}
		}
	}

	// Обновляем pass_id пользователя в базе данных
	err := database.SetUserPassID(username, nullablePassID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/gyms_passes", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/gyms_passes", http.StatusFound)
}

func SetGym(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.ServeFile(w, r, "./frontend/pages/not_auth_page.html")
		return
	}

	username := session.Values["username"].(string)

	gymIDStr := r.FormValue("gym_id")

	var nullableGymID sql.NullInt64
	if gymIDStr != "" {
		parsedGymID, err := strconv.Atoi(gymIDStr)
		if err == nil {
			nullableGymID = sql.NullInt64{Int64: int64(parsedGymID), Valid: true}
		}
	}

	// Обновляем pass_id пользователя в базе данных
	err := database.SetUserGymID(username, nullableGymID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/gyms", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/gyms", http.StatusFound)
}
