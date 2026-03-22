package http

import (
	"net/http"

	"example.com/webserver/internal/infrastructure/http/respond"
	"example.com/webserver/internal/infrastructure/i18n"
	"example.com/webserver/internal/modules/showcase/core/ports"
	"example.com/webserver/internal/modules/showcase/views"
	"example.com/webserver/internal/ui/layout"
)

type ShowcaseHandler struct {
	service ports.ShowcaseService
}

func New(svc ports.ShowcaseService) *ShowcaseHandler {
	return &ShowcaseHandler{service: svc}
}

func (h *ShowcaseHandler) Showcase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	items, err := h.service.GetShowcaseData(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	title := "Food Bridge"
	description := i18n.T(ctx, i18n.ShowcaseTitle)

	navItems := []layout.NavItem{
		{Label: i18n.T(ctx, i18n.ShowcaseTitle), Href: "/", Active: true},
	}
	user := layout.UserSummary{
		Name:  "Test User",
		Email: "test@example.com",
	}

	respond.View(w, r, views.ShowcasePage(title, description, navItems, user, items))
}
