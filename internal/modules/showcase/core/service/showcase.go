package service

import (
	"context"

	"example.com/webserver/internal/modules/showcase/core/domain"
	"example.com/webserver/internal/modules/showcase/core/ports"
)

type showcaseService struct {
	repo ports.ShowcaseRepository
}

func NewShowcaseService(repo ports.ShowcaseRepository) ports.ShowcaseService {
	return &showcaseService{repo: repo}
}

func (s *showcaseService) GetShowcaseData(ctx context.Context) ([]domain.ShowcaseItem, error) {
	return s.repo.FindAllShowcases(ctx)
}
