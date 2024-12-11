package services

import (
	"Gym-Service/internal/database"
	"fmt"
	"net/http"
	"strconv"
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
