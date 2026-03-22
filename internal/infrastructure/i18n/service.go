package i18n

import (
	"log/slog"

	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

// Service manages the application's internationalization, language resolution, and message translation.
type Service struct {
	logger           *slog.Logger
	config           Config
	translations     *catalog.Builder
	supportedTags    []language.Tag
	supportedOptions []LanguageOption
	matcher          language.Matcher
}

// New returns a new i18nService with a default config and the applied options.
func New(logger *slog.Logger, opts ...Option) *Service {
	cfg := defaultConfig()

	for _, opt := range opts {
		opt(&cfg)
	}

	service := Service{
		logger:       logger,
		config:       cfg,
		translations: catalog.NewBuilder(catalog.Fallback(DefaultTag)),
	}

	if err := service.loadTranslations(); err != nil {
		panic(err)
	}

	return &service
}
