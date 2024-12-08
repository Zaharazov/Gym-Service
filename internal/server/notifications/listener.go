package notifications

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func StartNotificationListener(connStr string) {
	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Создаем слушателя
	listener := pq.NewListener(connStr, 10, 90, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Fatalf("Listener error: %v", err)
		}
	})
	defer listener.Close()

	err = listener.Listen("user_registration")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	fmt.Println("Listening for user registration notifications...")

	// Обработка уведомлений
	for {
		select {
		case notification := <-listener.Notify:
			if notification != nil {
				fmt.Printf("Notification received: %s\n", notification.Extra)
			}
		}
	}
}
