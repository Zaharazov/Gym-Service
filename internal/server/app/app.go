package app

import (
	"Gym-Service/internal/database"
	"Gym-Service/internal/presentation/routers"
	"Gym-Service/internal/server/configs"
	"Gym-Service/internal/server/notifications"
	"database/sql"
	"log"
	"net/http"
)

func Run() {

	r := routers.NewRouter()

	log.Printf("Server started!")

	database.ConnectToPostrges()

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = database.CreateTrigger(db)
	if err != nil {
		log.Fatalf("Failed to create trigger: %v", err)
	}

	// Запускаем слушателя уведомлений в отдельной горутине
	go func() {
		notifications.StartNotificationListener(configs.DBPath)
	}()

	if err := http.ListenAndServe(configs.Port, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
