package i18n

import "time"

// Config holds the centralized settings for the i18n package.
type Config struct {
	CookieName   string
	CookiePath   string
	CookieDomain string
	CookieMaxAge time.Duration
	CookieSecure bool
}

// defaultConfig returns a Config with sensible default values.
func defaultConfig() Config {
	return Config{
		CookieName:   "lang",
		CookiePath:   "/",
		CookieMaxAge: 365 * 24 * time.Hour,
		CookieSecure: false,
	}
}

// Option defines a functional option for configuring the I18nService.
type Option func(*Config)

// WithCookieName sets the name of the language cookie.
func WithCookieName(name string) Option {
	return func(c *Config) {
		c.CookieName = name
	}
}

// WithCookiePath sets the path of the language cookie.
func WithCookiePath(path string) Option {
	return func(c *Config) {
		c.CookiePath = path
	}
}

// WithCookieDomain sets the domain of the language cookie.
func WithCookieDomain(domain string) Option {
	return func(c *Config) {
		c.CookieDomain = domain
	}
}

// WithCookieMaxAge sets the max age of the language cookie.
func WithCookieMaxAge(maxAge time.Duration) Option {
	return func(c *Config) {
		c.CookieMaxAge = maxAge
	}
}

// WithCookieSecure sets whether the language cookie is secure (HTTPS only).
func WithCookieSecure(secure bool) Option {
	return func(c *Config) {
		c.CookieSecure = secure
	}
}
