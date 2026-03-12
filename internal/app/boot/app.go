// Package boot handles application initialization.
package boot

import (
	"log/slog"
	"net/http"

	"example.com/webserver/internal/app/config"
	"example.com/webserver/internal/app/http/router"
	"example.com/webserver/internal/app/i18n"
)

// App represents the application.
type App struct {
	Logger      *slog.Logger
	I18nService *i18n.Service
	Router      http.Handler
}

// New creates a new application instance.
func New(logger *slog.Logger, cfg *config.Config) (*App, error) {
	i18nService := i18n.New(logger,
		i18n.WithCookieDomain(cfg.Domain),
		i18n.WithCookieSecure(cfg.IsProduction()),
	)
	httpRouter := router.New(logger, i18nService, cfg)

	return &App{
		Logger:      logger,
		I18nService: i18nService,
		Router:      httpRouter,
	}, nil
}
