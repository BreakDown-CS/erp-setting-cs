package helper

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func DateStringToDate(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date '%s': %w", value, err)
	}

	return t, nil
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return fallback
}
