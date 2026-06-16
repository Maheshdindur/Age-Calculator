package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBURL:      getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/user_db?sslmode=disable"),
		ServerPort: getEnv("SERVER_PORT", "3000"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
