// Package db holds the configuration for database connection setup and migrations.
package db

import (
	"context"
	"fmt"

	"example.com/webserver/internal/app/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool creates a new PostgreSQL connection pool.
func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dbCfg, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("unable to parse database url: %w", err)
	}

	// TODO: integrate ssl from config

	dbCfg.MaxConns = cfg.DBMaxConns
	dbCfg.MinConns = cfg.DBMinConns
	dbCfg.MaxConnIdleTime = cfg.DBMaxConnIdleTime
	dbCfg.MaxConnLifetime = cfg.DBMaxConnLifetime

	pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}
