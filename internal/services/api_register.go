package services

import (
	"Gym-Service/internal/database"
	"Gym-Service/internal/server/logger"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func RegUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirm_password := r.FormValue("confirm_password")
	age, _ := strconv.Atoi(r.FormValue("age"))
	sex := r.FormValue("sex")
	phone := r.FormValue("phone")

	if password != confirm_password {
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)

	err := database.AddUser(name, username, hashedPasswordStr, age, sex, phone)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	session, err := Store.Get(r, "session-name")
	if err != nil {
		logger.PublishLog("ERROR", fmt.Sprintf("Failed to register user: %v", err), "auth")
		log.Printf("Ошибка при получении сессии: %v", err) // лог в консоль
		http.Error(w, "Ошибка создания сессии", http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Values["role"] = "user"

	err = session.Save(r, w)
	if err != nil {
		log.Printf("Ошибка при сохранении сессии: %v", err) // лог в консоль
		http.Error(w, "Ошибка сохранения сессии", http.StatusInternalServerError)
		return
	}

	logger.PublishLog("INFO", fmt.Sprintf("User registered: %s", username), "auth")

	http.Redirect(w, r, "/gyms", http.StatusFound)
}
