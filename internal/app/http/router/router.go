// Package router provides the HTTP router.
package router

import (
	"log/slog"
	"net/http"

	"example.com/webserver/internal/app/config"
	"example.com/webserver/internal/app/db/store"
	"example.com/webserver/internal/app/http/middleware"
	"example.com/webserver/internal/app/i18n"
	homehttp "example.com/webserver/internal/modules/home/http"
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

	// Home
	homeHandler := homehttp.New(logger, storeQueries)
	mux.HandleFunc("GET /{$}", middleware.Chain(homeHandler.Index, middlewares...))

	return mux
}
