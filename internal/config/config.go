package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found, using system env")
	}

	return &Config{
		Port: getEnv("PORT", "3000"),

		DBHost:     getEnv("PG_DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("PG_DB_PORT", "5433"),
		DBUser:     getEnv("PG_DB_USERNAME", "postgres"),
		DBPassword: getEnv("PG_DB_PASSWORD", "1234"),
		DBName:     getEnv("PG_DB_NAME", "erp_cs"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
