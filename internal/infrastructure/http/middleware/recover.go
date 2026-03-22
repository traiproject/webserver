package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recover recovers from a panic within the application and returns a http 500.
func Recover(logger *slog.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(
						"recover from panic",
						slog.Any("error", err),
						slog.String("stack trace", string(debug.Stack())),
					)

					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next(w, r)
		}
	}
}
