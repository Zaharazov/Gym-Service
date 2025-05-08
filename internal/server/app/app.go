package app

import (
	"Gym-Service/internal/database"
	"Gym-Service/internal/presentation/routers"
	"Gym-Service/internal/server/configs"
	"Gym-Service/internal/server/notifications"
	"Gym-Service/internal/services"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/boj/redistore"
)

var store *redistore.RediStore

func Run() {

	var err error
	// 10 — max active connections; "tcp" and "localhost:6379" — адрес Redis
	store, err = redistore.NewRediStore(10, "tcp", "localhost:6379", "", "", []byte(configs.SESSION_KEY))
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	// Устанавливаем TTL для сессии
	store.SetMaxAge(int((2 * time.Minute).Seconds())) // 30 минут

	// Ключи в Redis будут автоматически удаляться
	store.SetKeyPrefix("session:")

	services.SetSessionStore(store)

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

	//Запускаем слушателя уведомлений в отдельной горутине
	go func() {
		notifications.StartNotificationListener(configs.DBPath)
	}()

	if err := http.ListenAndServe(configs.Port, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
