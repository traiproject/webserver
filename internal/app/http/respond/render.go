// Package respond provides helpers for HTTP responses.
package respond

import (
	"net/http"

	"github.com/a-h/templ"
)

// View renders a templ component to the response writer.
func View(w http.ResponseWriter, r *http.Request, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering view", http.StatusInternalServerError)
	}
}
