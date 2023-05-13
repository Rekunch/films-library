package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init(_ context.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file found")
	}

	return nil
}

func get(key string) string {
	return os.Getenv(key)
}
