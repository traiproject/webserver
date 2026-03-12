package middleware

import (
	"net/http"

	"example.com/webserver/internal/app/i18n"
)

// I18n adds the Localizer to the requests context.
func I18n(cfg i18n.ResolverConfig, service *i18n.Service) Middleware {
	if cfg.CookieName == "" {
		cfg.CookieName = "lang"
	}

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tag := service.ResolveTag(r, cfg)
			ctx := i18n.WithLocalizer(r.Context(), service.NewLocalizer(tag))
			next(w, r.WithContext(ctx))
		}
	}
}
