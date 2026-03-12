package i18n

import "golang.org/x/text/language"

// DefaultTag is the default language tag.
var DefaultTag = language.English

// LanguageOption represents a language option for the UI.
type LanguageOption struct {
	Tag      language.Tag
	Code     string
	LabelKey string
}

// BaseCode returns the base code of a language tag.
func BaseCode(tag language.Tag) string {
	base, _ := tag.Base()
	return base.String()
}

// MatchPreferred matches the preferred languages against the supported tags.
func (s *Service) MatchPreferred(preferred ...string) language.Tag {
	tag, _ := language.MatchStrings(s.matcher, preferred...)
	return tag
}

// ParseSupported parses a string into a supported language tag.
func (s *Service) ParseSupported(raw string) (language.Tag, bool) {
	tag, err := language.Parse(raw)
	if err != nil {
		return DefaultTag, false
	}

	matched, _, confidence := s.matcher.Match(tag)
	if confidence == language.No {
		return DefaultTag, false
	}

	return matched, true
}

// SupportedTags returns the supported tags.
func (s *Service) SupportedTags() []language.Tag {
	return s.supportedTags
}

// SupportedOptions returns the supported options.
func (s *Service) SupportedOptions() []LanguageOption {
	return s.supportedOptions
}
