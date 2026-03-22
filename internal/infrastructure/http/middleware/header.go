package middleware

import "net/http"

// SecurityHeaders adds security headers to the response.
func SecurityHeaders() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

			// TODO: Restrict img-src to specific trusted domains instead of allowing '*'
			w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' 'unsafe-inline'; "+
				"script-src 'self' 'unsafe-eval'; img-src * data:; base-uri 'self'; frame-ancestors 'none'")

			next(w, r)
		})
	}
}
