package i18n

import (
	"net/http"

	"golang.org/x/text/language"
)

// ResolverConfig is the configuration for the language resolver middleware.
type ResolverConfig struct {
	CookieName     string
	ResolveUserTag func(r *http.Request) (language.Tag, bool)
	ResolveURLTag  func(r *http.Request) (language.Tag, bool)
}

// ResolveTag resolves the language tag from the HTTP request.
func (s *Service) ResolveTag(r *http.Request, cfg ResolverConfig) language.Tag {
	if cfg.ResolveURLTag != nil {
		if tag, ok := cfg.ResolveURLTag(r); ok {
			return tag
		}
	}

	if cfg.ResolveUserTag != nil {
		if tag, ok := cfg.ResolveUserTag(r); ok {
			return tag
		}
	}

	if c, err := r.Cookie(cfg.CookieName); err == nil && c.Value != "" {
		if tag, ok := s.ParseSupported(c.Value); ok {
			return tag
		}
	}

	accept := r.Header.Get("Accept-Language")
	if accept != "" {
		tags, _, err := language.ParseAcceptLanguage(accept)
		if err == nil {
			matched, _, _ := s.matcher.Match(tags...)
			return matched
		}
	}

	return DefaultTag
}
