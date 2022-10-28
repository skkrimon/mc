package util

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env.local")

	if err == nil {
		return
	}

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env file.")
	}
}
