package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port     = GetConfigs("httpPort")
	DBPath   = GetConfigs("dbPath")
	DBName   = GetConfigs("dbName")
	Username = GetConfigs("username")
	Password = GetConfigs("password")
)

func GetConfigs(param string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Data, exists := os.LookupEnv(param)

	if exists {
		log.Printf("%s is %s", param, Data)
	} else {
		log.Fatalf("%s is missing", param)
	}

	return Data
}
