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

func CreateTrigger(db *sql.DB) error {

	checkTriggerQuery := `
	SELECT EXISTS (
		SELECT 1 
		FROM pg_trigger 
		WHERE tgname = 'user_registration_trigger' 
	);`

	var exists bool
	err := db.QueryRow(checkTriggerQuery).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if trigger exists: %v", err)
		return err
	}

	if exists {
		log.Println("Trigger already exists.")
		return nil
	}

	// SQL для создания функции
	createFunction := `
	CREATE OR REPLACE FUNCTION log_user_registration()
	RETURNS TRIGGER AS $$
	BEGIN
		PERFORM pg_notify('user_registration', 'User registered with ID: ' || NEW.user_id || ', Name: ' || NEW.name);
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	`

	// SQL для создания триггера
	createTrigger := `
	CREATE TRIGGER user_registration_trigger
	AFTER INSERT ON users
	FOR EACH ROW
	EXECUTE FUNCTION log_user_registration();
	`

	// Выполняем запросы
	_, err = db.Exec(createFunction)
	if err != nil {
		log.Printf("Error creating function: %v", err)
		return err
	}

	_, err = db.Exec(createTrigger)
	if err != nil {
		log.Printf("Error creating trigger: %v", err)
		return err
	}

	log.Println("Trigger and function created successfully.")
	return nil
}
