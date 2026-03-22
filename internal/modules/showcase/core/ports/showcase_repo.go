package ports

import (
	"context"

	"example.com/webserver/internal/modules/showcase/core/domain"
)

// ShowcaseRepository defines the data persistence requirements.
type ShowcaseRepository interface {
	FindAllShowcases(ctx context.Context) ([]domain.ShowcaseItem, error)
}
