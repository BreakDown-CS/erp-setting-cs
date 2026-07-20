package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/BreakDown-CS/erp-setting-cs/internal/config"
	"github.com/BreakDown-CS/erp-setting-cs/internal/helper"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

func ConnPostgres(cfg *config.Config) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	configPool, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	configPool.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   &PgxLogger{},
		LogLevel: tracelog.LogLevelInfo,
	}

	db, err := pgxpool.NewWithConfig(context.Background(), configPool)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return db, nil
}

type PgxLogger struct{}

func (l *PgxLogger) Log(
	ctx context.Context,
	level tracelog.LogLevel,
	msg string,
	data map[string]any,
) {

	if msg != "Query" {
		return
	}

	sql, ok := data["sql"].(string)
	if !ok {
		return
	}

	switch strings.ToLower(strings.TrimSpace(sql)) {
	case "begin", "commit", "rollback":
		return
	}

	args, _ := data["args"].([]any)

	helper.SQL(
		sql,
		args,
		data["time"],
	)
}
