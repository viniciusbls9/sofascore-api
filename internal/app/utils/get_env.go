package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func HandlerGetEnv() (string, string) {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env")
		return "", ""
	}

	if dbName == "" {
		log.Fatal("DB_NAME is not found in the env")
		return "", ""
	}

	return dbURL, dbName
}
