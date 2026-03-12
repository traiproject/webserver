package i18n

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

// Localizer is a thread-safe helper for localizing strings.
type Localizer struct {
	tag          language.Tag
	printer      *message.Printer
	translations *catalog.Builder
	options      []LanguageOption
}

// NewLocalizer creates a new localizer for the given language tag.
func (s *Service) NewLocalizer(tag language.Tag) *Localizer {
	if _, ok := s.ParseSupported(tag.String()); !ok {
		tag = DefaultTag
	}

	return &Localizer{
		tag:          tag,
		printer:      message.NewPrinter(tag, message.Catalog(s.translations)),
		translations: s.translations,
		options:      s.SupportedOptions(),
	}
}

// Sprintf formats a string with the localizer's language.
func (l *Localizer) Sprintf(key string, args ...any) string {
	if l == nil {
		if len(args) == 0 {
			return key
		}
		return fmt.Sprintf(key, args...)
	}
	return l.printer.Sprintf(key, args...)
}

// Sprint formats a string with the localizer's language.
func (l *Localizer) Sprint(v ...any) string {
	if l == nil {
		return fmt.Sprint(v...)
	}
	return l.printer.Sprint(v...)
}
