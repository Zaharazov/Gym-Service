package services

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// добавить логику для бд !!!!!!!!!!!!!!!!!!!!!!!!!
var users = map[string]string{
	"user1": "password1", // Логин -> Пароль
	"user2": "password2",
}

var admins = map[string]string{
	"admin": "adminpass", // Логин -> Пароль
}

var store = sessions.NewCookieStore([]byte("super-secret-key"))

// добавить логику для бд !!!!!!!!!!!!!!!!!!!!!!!!!
func authenticate(username, password string) (string, bool) {
	// Проверяем в таблице пользователей
	if pass, exists := users[username]; exists && pass == password {
		return "user", true
	}

	// Проверяем в таблице администраторов
	if pass, exists := admins[username]; exists && pass == password {
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
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
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
