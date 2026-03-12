package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader intercepts the status code before passing it to the real ResponseWriter.
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logging logs the request once completed.
func Logging(logger *slog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrappedWriter := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			next(wrappedWriter, r)

			logger.Info("request completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("response_code", wrappedWriter.status),
				slog.Duration("duration", time.Since(start)),
			)
		}
	}
}
