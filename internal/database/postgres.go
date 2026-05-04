package database

import (
	"context"
	"fmt"

	"github.com/BreakDown-CS/erp-setting-cs/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnPostgres(cfg *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBstaff,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic("cannot connect db")
	}

	return db
}
