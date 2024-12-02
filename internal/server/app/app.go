package app

import (
	"Gym-Service/internal/database"
	"Gym-Service/internal/presentation/routers"
	"Gym-Service/internal/server/configs"
	"fmt"
	"log"
	"net/http"
)

func Run() {

	r := routers.NewRouter()

	log.Printf("Server started!")

	database.ConnectToPostrges()

	if err := http.ListenAndServe(configs.Port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
