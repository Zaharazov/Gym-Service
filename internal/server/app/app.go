package app

import (
	"Gym-Service/internal/database"
	"Gym-Service/internal/presentation/routers"
	"Gym-Service/internal/server/configs"
	"Gym-Service/internal/server/notifications"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func Run() {

	r := routers.NewRouter()

	log.Printf("Server started!")

	database.ConnectToPostrges()

	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)

	err = database.CreateTrigger(db)
	if err != nil {
		log.Fatalf("Failed to create trigger: %v", err)
	}

	// Запускаем слушателя уведомлений в отдельной горутине
	go func() {
		notifications.StartNotificationListener(configs.DBPath)
	}()

	if err := http.ListenAndServe(configs.Port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
