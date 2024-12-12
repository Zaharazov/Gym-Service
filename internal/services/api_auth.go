package services

import (
	"Gym-Service/internal/database"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

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
	if !authenticated {

		// Сохраняем ошибку в сессии
		session, _ := store.Get(r, "session-name")
		session.Values["error"] = "Неверный логин или пароль"
		session.Values["username"] = username // Сохраняем введенное имя пользователя
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Успешный вход, сохраняем данные в сессию
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Values["role"] = role
	session.Save(r, w)

	http.Redirect(w, r, "/gyms", http.StatusFound)
}

func EndSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Values["username"] = nil
	session.Values["role"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}
