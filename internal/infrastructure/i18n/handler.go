package i18n

import (
	"net/http"
)

const (
	// MaxFormSize is the maximum size of the form data (2MB).
	MaxFormSize = 2 * 1024 * 1024
)

// LanguageHandler creates a handler that sets the language cookie using the service config.
func (s *Service) LanguageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, MaxFormSize)
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		tag, ok := s.ParseSupported(r.FormValue("lang"))
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     s.config.CookieName,
			Value:    BaseCode(tag), // Assumes BaseCode is defined in your package
			Path:     s.config.CookiePath,
			MaxAge:   int(s.config.CookieMaxAge.Seconds()),
			HttpOnly: true,
			Domain:   s.config.CookieDomain,
			Secure:   s.config.CookieSecure,
			SameSite: http.SameSiteLaxMode,
		})

		redirectTo := r.FormValue("redirect_to")
		if redirectTo == "" {
			redirectTo = r.Referer()
		}
		if redirectTo == "" {
			redirectTo = "/"
		}

		http.Redirect(w, r, redirectTo, http.StatusSeeOther)
	}
}
