// Package router provides the HTTP router.
package router

import (
	"log/slog"
	"net/http"

	"example.com/webserver/internal/app/config"
	"example.com/webserver/internal/infrastructure/db/store"
	"example.com/webserver/internal/infrastructure/http/middleware"
	"example.com/webserver/internal/infrastructure/i18n"
	showcaserepo "example.com/webserver/internal/modules/showcase/adapters/postgres"
	showcaseservice "example.com/webserver/internal/modules/showcase/core/service"
	showcasehttp "example.com/webserver/internal/modules/showcase/http"
)

// New creates a new HTTP router.
func New(logger *slog.Logger, i18nService *i18n.Service, cfg *config.Config, storeQueries *store.Queries) http.Handler {
	middlewares := []middleware.Middleware{
		middleware.Logging(logger),
		middleware.Recover(logger),
		middleware.SecurityHeaders(),
		middleware.I18n(i18n.ResolverConfig{}, i18nService),
	}

	mux := http.NewServeMux()

	serveStatic(mux, cfg)

	mux.Handle("POST /language", i18nService.LanguageHandler())

	// Showcase
	showcaseRepo := showcaserepo.NewShowcaseRepository(storeQueries)
	showcaseService := showcaseservice.NewShowcaseService(showcaseRepo)
	showcaseHandler := showcasehttp.New(showcaseService)
	mux.HandleFunc("GET /{$}", middleware.Chain(showcaseHandler.Showcase, middlewares...))

	return mux
}
