package ports

import (
	"context"

	"example.com/webserver/internal/modules/showcase/core/domain"
)

// ShowcaseService defines the high-level business actions for the showcase module.
type ShowcaseService interface {
	GetShowcaseData(ctx context.Context) ([]domain.ShowcaseItem, error)
}
