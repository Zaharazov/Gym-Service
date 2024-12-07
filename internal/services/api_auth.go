package services

import (
	"Gym-Service/internal/database"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func authenticate(username, password string) (string, bool) {
	// Проверяем в таблице пользователей

	exist, pass, err := database.FindUser(username)

	if err != nil {
		fmt.Printf("Ошибка при проверке пользователя: %v\n", err)
		return "", false
	}

	if exist && pass == password {
		return "user", true
	}

	// Проверяем в таблице тренеров

	exist, pass, err = database.FindCoach(username)

	if err != nil {
		fmt.Printf("Ошибка при проверке тренера: %v\n", err)
		return "", false
	}

	if exist && pass == password {
		return "coach", true
	}

	// Проверяем в таблице администраторов

	exist, pass, err = database.FindAdmin(username)

	if err != nil {
		fmt.Printf("Ошибка при проверке админа: %v\n", err)
		return "", false
	}

	if exist && pass == password {
		return "admin", true
	}

	// Если логин/пароль не найден
	return "", false
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Аутентифицируем пользователя
	role, authenticated := authenticate(username, password)
	if !authenticated {
		http.Redirect(w, r, "/", http.StatusFound)
		//http.Error(w, "Invalid username or password", http.StatusUnauthorized)
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
