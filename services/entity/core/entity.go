package core

import (
	"context"

	"github.com/seboste/ai-agent-test/services/entity/ports"
)

type EntityService struct {
	repo     ports.Repo
	notifier ports.Notifier
}

func NewEntityService(repo ports.Repo, notifier ports.Notifier) *EntityService {
	return &EntityService{
		repo:     repo,
		notifier: notifier,
	}
}

var _ ports.Api = (*EntityService)(nil)

func (s *EntityService) Set(entity ports.Entity, ctx context.Context) error {
	err := s.repo.Store(entity, ctx)
	if err != nil {
		return err
	}
	s.notifier.EntityChanged(entity, ctx)
	return nil
}

func (s *EntityService) Get(id string, ctx context.Context) (ports.Entity, error) {
	entity, err := s.repo.FindById(id, ctx)
	if err != nil {
		return ports.Entity{}, err
	}
	if entity.Id != id {
		return ports.Entity{}, ports.ErrEntityNotFound
	}
	return entity, nil
}
