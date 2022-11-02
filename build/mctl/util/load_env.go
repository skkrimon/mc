package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	err := godotenv.Load(".env.local")

	if err != nil && (os.Getenv(key) != "") {
		return os.Getenv(key)
	}

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("env var not found")
	}

	return os.Getenv(key)
}
