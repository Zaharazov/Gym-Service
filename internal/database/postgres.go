package database

import (
	"Gym-Service/internal/server/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToPostrges() {
	connStr := configs.DBPath
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
