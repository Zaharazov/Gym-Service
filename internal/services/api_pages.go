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

	gyms, err := database.FetchGyms()
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
		//fmt.Println(err)
		http.Error(w, "Failed to fetch gym details", http.StatusInternalServerError)
		return
	}

	// Получаем данные о тренажерах в этом зале
	equipments, err := database.FetchEquipmentsForGym(gymID)
	if err != nil {

		http.Error(w, "Failed to fetch equipments", http.StatusInternalServerError)
		return
	}

	// Получаем данные о групповых тренировках
	events, err := database.FetchEventsForGym(gymID)
	if err != nil {
		//fmt.Println(err)
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}

	// Получаем данные о тренерах
	coaches, err := database.FetchTrainersForGym(gymID)
	if err != nil {
		http.Error(w, "Failed to fetch coaches", http.StatusInternalServerError)
		return
	}

	user_role := session.Values["role"].(string)

	// Подготавливаем данные для шаблона
	data := struct {
		Role       string
		Gym        models.Gym
		Coaches    []models.Coach
		Equipments []models.Equipment
		Events     []models.Event
	}{
		Role:       user_role,
		Gym:        gym,
		Coaches:    coaches,
		Equipments: equipments,
		Events:     events,
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

func GetGymsPassesPage(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	gyms_passes, err := database.FetchGymsPasses()
	if err != nil {
		http.Error(w, "Failed to fetch gyms passes", http.StatusInternalServerError)
		return
	}

	user_role := session.Values["role"].(string)

	// Подготавливаем данные для шаблона
	data := struct {
		Role   string
		Passes []models.Pass
	}{
		Role:   user_role,
		Passes: gyms_passes,
	}

	tmpl, err := template.ParseFiles("./frontend/pages/gyms_passes_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func GetEquipmentPage(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	equipments, err := database.FetchEquipment()
	if err != nil {
		http.Error(w, "Failed to fetch equipment", http.StatusInternalServerError)
		return
	}

	user_role := session.Values["role"].(string)

	// Подготавливаем данные для шаблона
	data := struct {
		Role      string
		Equipment []models.Equipment
	}{
		Role:      user_role,
		Equipment: equipments,
	}

	tmpl, err := template.ParseFiles("./frontend/pages/equipment_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func GetAdminPage(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	auth, ok := session.Values["authenticated"].(bool)
	user_role := session.Values["role"].(string)

	if !ok || !auth || user_role != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		gym         models.Gym
		users       []models.User
		coaches     []models.Coach
		admins      []models.Admin
		gyms        []models.Gym
		gyms_passes []models.Pass
		equipments  []models.Equipment
		events      []models.Event
		err         error
	)

	gymID := r.FormValue("gym_id")

	if gymID != "" {

		users, err = database.FetchUsersForGym(gymID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}

		coaches, err = database.FetchTrainersForGym(gymID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch coaches", http.StatusInternalServerError)
			return
		}

		admins, err = database.FetchAdminsForGym(gymID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch admins", http.StatusInternalServerError)
			return
		}

		gym, err = database.FetchGymByID(gymID)
		if err != nil {
			fmt.Println(err)
		} else {
			gyms = append(gyms, gym)
		}

		gyms_passes, err = database.FetchGymsPasses()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch gyms passes", http.StatusInternalServerError)
			return
		}

		equipments, err = database.FetchEquipmentsForGym(gymID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch equipment", http.StatusInternalServerError)
			return
		}

		events, err = database.FetchEventsForGym(gymID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
			return
		}
	} else {
		users, err = database.FetchUsers()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		coaches, err = database.FetchCoaches()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch coaches", http.StatusInternalServerError)
			return
		}
		admins, err = database.FetchAdmins()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch admins", http.StatusInternalServerError)
			return
		}
		gyms, err = database.FetchGyms()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch gyms", http.StatusInternalServerError)
			return
		}
		gyms_passes, err = database.FetchGymsPasses()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch gyms passes", http.StatusInternalServerError)
			return
		}
		equipments, err = database.FetchEquipment()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch equipment", http.StatusInternalServerError)
			return
		}
		events, err = database.FetchEvents()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
			return
		}
	}

	// Подготавливаем данные для шаблона
	data := struct {
		Role      string
		Users     []models.User
		Coaches   []models.Coach
		Admins    []models.Admin
		Gyms      []models.Gym
		Passes    []models.Pass
		Equipment []models.Equipment
		Events    []models.Event
	}{
		Role:      user_role,
		Users:     users,
		Coaches:   coaches,
		Admins:    admins,
		Gyms:      gyms,
		Passes:    gyms_passes,
		Equipment: equipments,
		Events:    events,
	}

	tmpl, err := template.ParseFiles("./frontend/pages/admin_page.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
