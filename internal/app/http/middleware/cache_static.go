package middleware

import "net/http"

// CacheStatic adds caching for static resources in production environment.
func CacheStatic(isProd bool) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if isProd {
				// Cache aggressively in production (1 year)
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			} else {
				// Force browser to fetch fresh files during development
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			}

			next(w, r)
		}
	}
}
