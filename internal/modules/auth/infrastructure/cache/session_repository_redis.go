package cache

import (
	"context"
	"user-service/internal/modules/auth/domain"
	"user-service/internal/shared/components"
)

var _ domain.SessionRepositoryPort = &sessionRepository{}

type sessionRepository struct {
	redis *components.Redis
}

func NewSessionRepository(redis *components.Redis) *sessionRepository {
	return &sessionRepository{redis: redis}
}

func (s *sessionRepository) CreateSession(ctx context.Context, session *domain.Session) error {
	return nil
}

func (s *sessionRepository) GetSession(ctx context.Context, sessionID string) (*domain.Session, error) {
	return nil, nil
}

func (s *sessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return nil
}
