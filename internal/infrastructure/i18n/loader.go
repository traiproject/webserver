package i18n

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"strings"

	"github.com/goccy/go-yaml"
	"golang.org/x/text/language"
)

//go:embed translations/*.yaml
var translationsFs embed.FS

const translationsDirName = "translations"

func (s *Service) loadTranslations() error {
	translationFiles, err := fs.ReadDir(translationsFs, translationsDirName)
	if err != nil {
		return fmt.Errorf("could not read locales directory: %w", err)
	}

	var tags []language.Tag
	var options []LanguageOption

	for _, file := range translationFiles {

		code := strings.TrimSuffix(file.Name(), ".yaml")

		tag, err := language.Parse(code)
		if err != nil {
			s.logger.Warn("invalid language tag in filename", slog.String("file", file.Name()), slog.Any("error", err))
			continue
		}

		tags = append(tags, tag)
		options = append(options, LanguageOption{
			Tag:      tag,
			Code:     code,
			LabelKey: "lang." + code,
		})

		data, err := translationsFs.ReadFile(fmt.Sprintf("%s/%s", translationsDirName, file.Name()))
		if err != nil {
			return fmt.Errorf("could not read locales file: %w", err)
		}

		translationsMap := make(map[string]string)
		if err := yaml.Unmarshal([]byte(data), &translationsMap); err != nil {
			return fmt.Errorf("could not unmarshal locales file: %w", err)
		}

		for k, v := range translationsMap {
			if err := s.translations.SetString(tag, k, v); err != nil {
				s.logger.Error("could not set translation", slog.String("key", k), slog.String("value", v))
			}
		}
	}

	if len(tags) == 0 {
		tags = []language.Tag{DefaultTag}
	}

	s.supportedTags = tags
	s.supportedOptions = options
	s.matcher = language.NewMatcher(tags)

	return nil
}
