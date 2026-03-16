// Package boot handles application initialization.
package boot

import (
	"context"
	"log/slog"
	"net/http"

	"example.com/webserver/internal/app/config"
	"example.com/webserver/internal/app/db"
	"example.com/webserver/internal/app/db/store"
	"example.com/webserver/internal/app/http/router"
	"example.com/webserver/internal/app/i18n"
	"github.com/jackc/pgx/v5/pgxpool"
)

// App represents the application.
type App struct {
	Logger      *slog.Logger
	I18nService *i18n.Service
	Router      http.Handler
	DBPool      *pgxpool.Pool
	Store       *store.Queries
}

// New creates a new application instance.
func New(ctx context.Context, logger *slog.Logger, cfg *config.Config) (*App, func(), error) {
	dbPool, dbErr := db.NewPool(ctx, cfg)
	if dbErr != nil {
		return nil, nil, dbErr
	}

	migErr := db.RunMigrations(ctx, dbPool, logger)
	if migErr != nil {
		return nil, nil, migErr
	}

	seedErr := db.RunSeeds(ctx, dbPool, logger, cfg)
	if seedErr != nil {
		return nil, nil, seedErr
	}

	storeQueries := store.New(dbPool)

	cleanup := func() {
		logger.Info("closing database connection pool")
		dbPool.Close()
	}

	i18nService := i18n.New(logger,
		i18n.WithCookieDomain(cfg.Domain),
		i18n.WithCookieSecure(cfg.IsProduction()),
	)

	httpRouter := router.New(logger, i18nService, cfg, storeQueries)

	return &App{
		Logger:      logger,
		I18nService: i18nService,
		Router:      httpRouter,
		DBPool:      dbPool,
		Store:       storeQueries,
	}, cleanup, nil
}
