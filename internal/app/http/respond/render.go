// Package respond provides helpers for HTTP responses.
package respond

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

// View renders a templ component to the response writer.
// View expects all data to be fetched in advance by the handler.
// It cannot propagate render errors to the user properly in favor of performance.
func View(w http.ResponseWriter, r *http.Request, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		// Just log it. The only way it fails here is a network disconnect
		// or a catastrophic templ bug, in which case we can't fix the HTTP response anyway.
		slog.ErrorContext(r.Context(), "failed to write templ component to response", slog.Any("error", err), slog.String("path", r.URL.Path))
	}
}
