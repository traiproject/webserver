package postgres

import (
	"context"

	"example.com/webserver/internal/infrastructure/db/store"
	"example.com/webserver/internal/modules/showcase/core/domain"
	"example.com/webserver/internal/modules/showcase/core/ports"
)

type showcaseRepo struct {
	q *store.Queries
}

func NewShowcaseRepository(q *store.Queries) ports.ShowcaseRepository {
	return &showcaseRepo{q: q}
}

func (r *showcaseRepo) FindAllShowcases(ctx context.Context) ([]domain.ShowcaseItem, error) {
	users, err := r.q.ListShowcaseItem(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]domain.ShowcaseItem, len(users))
	for i, u := range users {
		items[i] = domain.ShowcaseItem{
			ID:    u.ID,
			Title: u.Title,
		}
	}
	return items, nil
}
