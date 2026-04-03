package service

import (
	"context"
	"shortner/internal/database/repository"
	db "shortner/internal/database/sqlc"
)

type UserAgentService struct {
	repo repository.UserAgentRepository
}

// NewUserAgentService - конструктор для UserAgentService
func NewUserAgentService(repo repository.UserAgentRepository) *UserAgentService {
	return &UserAgentService{repo: repo}
}

// CreateUserAgent - создает новый user_agent в базе данных
func (s *UserAgentService) CreateUserAgent(ctx context.Context, fingerprintID int64, userAgent string) (*db.UserAgent, error) {
	return s.repo.CreateUserAgent(ctx, fingerprintID, userAgent)
}

// GetUserAgents - возвращает все user_agents из базы данных
func (s *UserAgentService) GetUserAgents(ctx context.Context) ([]*db.UserAgent, error) {
	return s.repo.GetUserAgents(ctx)
}

// GetUserAgentById - возвращает user_agent по его ID
func (s *UserAgentService) GetUserAgentById(ctx context.Context, id int64) (*db.UserAgent, error) {
	return s.repo.GetUserAgentById(ctx, id)
}

// GetUserAgentsByFingerprintId - возвращает все user_agents по ID fingerprint
func (s *UserAgentService) GetUserAgentsByFingerprintId(ctx context.Context, fingerprintID int64) ([]*db.UserAgent, error) {
	return s.repo.GetUserAgentsByFingerprintId(ctx, fingerprintID)
}

// UpdateUserAgentLastAccessedById - обновляет время последнего доступа user_agent по его ID
func (s *UserAgentService) UpdateUserAgentLastAccessedById(ctx context.Context, id int64) (*db.UserAgent, error) {
	return s.repo.UpdateUserAgentLastAccessedById(ctx, id)
}

// GetUserAgentByFpIdAgent - возвращает user_agent по ID fingerprint и user_agent
func (s *UserAgentService) GetUserAgentByFpIdAgent(ctx context.Context, fingerprintID int64, userAgent string) (*db.UserAgent, error) {
	return s.repo.GetUserAgentByFpIdAgent(ctx, fingerprintID, userAgent)
}
