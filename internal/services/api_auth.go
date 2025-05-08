package services

import (
	"Gym-Service/internal/database"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Store sessions.Store // глобальный store, который будет установлен из main.go

func SetSessionStore(s sessions.Store) {
	Store = s
}

func authenticate(username, password string) (string, bool) {
	// Проверяем в таблице пользователей

	exist, hashPass, err := database.FindUser(username)

	if err != nil {
		log.Printf("Ошибка при проверке пользователя %s: %v", username, err)
		return "", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if exist && err == nil {
		return "user", true
	}

	// Проверяем в таблице тренеров

	exist, hashPass, err = database.FindCoach(username)

	if err != nil {
		log.Printf("Ошибка при проверке тренера %s: %v", username, err)
		return "", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if exist && err == nil {
		return "coach", true
	}

	// Проверяем в таблице администраторов

	exist, hashPass, err = database.FindAdmin(username)

	if err != nil {
		log.Printf("Ошибка при проверке админа %s: %v", username, err)
		return "", false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	if exist && err == nil {
		return "admin", true
	}

	// Если логин/пароль не найден
	log.Printf("Неверный пароль: %v", err)
	return "", false
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Аутентифицируем пользователя
	role, authenticated := authenticate(username, password)

	session, err := Store.Get(r, "session-name")
	if err != nil {
		log.Printf("Ошибка при получении сессии: %v", err) // лог в консоль
		http.Error(w, "Ошибка получения сессии", http.StatusInternalServerError)
		return
	}

	if !authenticated {
		session.Values["error"] = "Неверный логин или пароль"
		session.Values["username"] = username
		if err := session.Save(r, w); err != nil {
			log.Printf("Ошибка сохранения сессии: %v", err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Values["role"] = role
	if err := session.Save(r, w); err != nil {
		log.Printf("Ошибка сохранения сессии: %v", err)
	}

	http.Redirect(w, r, "/gyms", http.StatusFound)
}

func EndSession(w http.ResponseWriter, r *http.Request) {

	session, err := Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Ошибка завершения сессии", http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = false
	session.Values["username"] = nil
	session.Values["role"] = nil

	if err := session.Save(r, w); err != nil {
		log.Printf("Ошибка сохранения сессии: %v", err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
