package config

import (
	"log"

	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
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
		Port: helper.GetEnv("PORT", "3000"),

		DBHost:     helper.GetEnv("PG_DB_HOST", "127.0.0.1"),
		DBPort:     helper.GetEnv("PG_DB_PORT", "5433"),
		DBUser:     helper.GetEnv("PG_DB_USERNAME", "postgres"),
		DBPassword: helper.GetEnv("PG_DB_PASSWORD", "1234"),
		DBName:     helper.GetEnv("PG_DB_NAME", "erp_cs"),
		DBSSLMode:  helper.GetEnv("DB_SSLMODE", "disable"),
	}
}
