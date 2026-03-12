// Package http provides HTTP handlers for the home module.
package http

import (
	"log/slog"
	"net/http"

	"example.com/webserver/internal/app/http/respond"
	"example.com/webserver/internal/app/i18n"
	"example.com/webserver/internal/modules/home/views"
	"example.com/webserver/internal/ui/layout"
)

// HomeHandler handles home page requests.
type HomeHandler struct {
	logger *slog.Logger
}

// New creates a new HomeHandler.
func New(logger *slog.Logger) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

// Index renders the home page.
func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Grab context for i18n
	ctx := r.Context()

	title := "Food Bridge"
	description := i18n.T(ctx, i18n.HomeTitle)

	navItems := []layout.NavItem{
		{Label: i18n.T(ctx, i18n.HomeTitle), Href: "/", Active: true},
	}
	user := layout.UserSummary{
		Name:  "Test User",
		Email: "test@example.com",
	}

	respond.View(w, r, views.IndexPage(title, description, navItems, user))
}
