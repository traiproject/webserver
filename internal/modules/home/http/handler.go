// Package http provides HTTP handlers for the home module.
package http

import (
	"log/slog"
	"net/http"

	"example.com/webserver/internal/app/db/store"
	"example.com/webserver/internal/app/http/respond"
	"example.com/webserver/internal/app/i18n"
	"example.com/webserver/internal/modules/home/views"
	"example.com/webserver/internal/ui/layout"
)

// HomeHandler handles home page requests.
type HomeHandler struct {
	logger *slog.Logger
	store  *store.Queries
}

// New creates a new HomeHandler.
func New(logger *slog.Logger, storeQueries *store.Queries) *HomeHandler {
	return &HomeHandler{
		logger: logger,
		store:  storeQueries,
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

	users, err := h.store.ListUsers(ctx)
	if err != nil {
		http.Error(w, "could not fetch users", http.StatusInternalServerError)
		return
	}

	vm := views.IndexVM{
		Users: users,
	}

	respond.View(w, r, views.IndexPage(title, description, navItems, user, vm))
}
