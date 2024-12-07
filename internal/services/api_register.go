package services

import (
	"Gym-Service/internal/database"
	"fmt"
	"net/http"
)

func RegUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirm_password := r.FormValue("confirm_password")

	if password != confirm_password {
		http.Redirect(w, r, "/register", http.StatusFound)
		//http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	err := database.AddUser(name, username, password)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/register", http.StatusFound)
		//http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Values["role"] = "user"
	session.Save(r, w)

	http.Redirect(w, r, "/gyms", http.StatusFound)
}
