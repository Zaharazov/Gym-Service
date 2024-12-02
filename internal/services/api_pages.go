package services

import (
	"Gym-Service/internal/database"
	"Gym-Service/internal/domain/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

func loadHTMLFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return string(content)
}

func GetAuthPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := loadHTMLFile("./frontend/pages/auth_page.html")

	fmt.Fprint(w, html)
}

func GetGymsPage(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	gyms, err := database.FetchGymsFromDB()
	if err != nil {
		http.Error(w, "Failed to fetch gyms", http.StatusInternalServerError)
		return
	}

	user_role := session.Values["role"].(string)

	// Подготавливаем данные для шаблона
	data := struct {
		Role string
		Gyms []models.Gym
	}{
		Role: user_role,
		Gyms: gyms,
	}

	tmpl, err := template.ParseFiles("./frontend/pages/gyms_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func GetGymPageById(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Извлечение gym_id из URL
	vars := mux.Vars(r)
	gymID := vars["gym_id"]

	// Получаем данные о конкретном зале
	gym, err := database.FetchGymByID(gymID)
	if err != nil {
		http.Error(w, "Failed to fetch gym details", http.StatusInternalServerError)
		return
	}

	// Получаем данные о тренажерах в этом зале

	// Получаем данные о групповых тренировках

	// Получаем данные о тренерах
	coaches, err := database.FetchTrainersForGym(gymID)
	if err != nil {
		http.Error(w, "Failed to fetch coaches", http.StatusInternalServerError)
		return
	}

	user_role := session.Values["role"].(string)

	// Подготавливаем данные для шаблона
	data := struct {
		Role    string
		Gym     models.Gym
		Coaches []models.Coach
	}{
		Role:    user_role,
		Gym:     gym,
		Coaches: coaches,
	}

	// Загружаем HTML-шаблон
	tmpl, err := template.ParseFiles("./frontend/pages/gym_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	// Рендерим HTML с данными
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
