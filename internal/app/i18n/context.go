// Package i18n provides internationalization support.
package i18n

import (
	"context"

	"golang.org/x/text/language"
)

type contextKey struct{}

var localizerContextKey = contextKey{}

// WithLocalizer returns a new context with the given localizer.
func WithLocalizer(ctx context.Context, l *Localizer) context.Context {
	return context.WithValue(ctx, localizerContextKey, l)
}

// FromContext returns the localizer from the context.
func FromContext(ctx context.Context) *Localizer {
	if l, ok := ctx.Value(localizerContextKey).(*Localizer); ok && l != nil {
		return l
	}
	return nil
}

// SupportedOptions returns the available language options from the context.
func SupportedOptions(ctx context.Context) []LanguageOption {
	localizer := FromContext(ctx)
	if localizer == nil {
		return nil
	}
	return localizer.options
}

// CurrentTag returns the current language tag from the context.
func CurrentTag(ctx context.Context) language.Tag {
	return FromContext(ctx).tag
}

// T translates a key with arguments using the context's localizer.
func T(ctx context.Context, key string, args ...any) string {
	localizer := FromContext(ctx)
	if localizer == nil {
		return key
	}
	return localizer.Sprintf(key, args...)
}
