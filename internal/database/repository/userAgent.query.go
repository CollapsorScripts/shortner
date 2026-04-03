package repository

import (
	"context"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userAgentRepo struct {
	db      *pgxpool.Pool
	cfg     *config.Config
	queries *db.Queries
}

// NewUserAgentRepository - создает новый репозиторий для работы с user_agent
func NewUserAgentRepository(conn *pgxpool.Pool, cfg *config.Config) UserAgentRepository {
	return &userAgentRepo{
		db:      conn,
		cfg:     cfg,
		queries: db.New(conn),
	}
}

// CreateUserAgent - создает новый user_agent в базе данных
func (r *userAgentRepo) CreateUserAgent(ctx context.Context, fingerprintID int64, userAgent string) (*db.UserAgent, error) {
	return r.queries.CreateUserAgent(ctx, db.CreateUserAgentParams{
		FingerprintID: fingerprintID,
		Agent:         userAgent,
	})
}

// GetUserAgents - возвращает все user_agents из базы данных
func (r *userAgentRepo) GetUserAgents(ctx context.Context) ([]*db.UserAgent, error) {
	return r.queries.GetUserAgents(ctx)
}

// GetUserAgentById - возвращает user_agent по его ID
func (r *userAgentRepo) GetUserAgentById(ctx context.Context, id int64) (*db.UserAgent, error) {
	return r.queries.GetUserAgentById(ctx, id)
}

// GetUserAgentsByFingerprintId - возвращает все user_agents по ID fingerprint
func (r *userAgentRepo) GetUserAgentsByFingerprintId(ctx context.Context, fingerprintID int64) ([]*db.UserAgent, error) {
	return r.queries.GetUserAgentsByFingerprintId(ctx, fingerprintID)
}

// UpdateUserAgentLastAccessedById - обновляет время последнего доступа user_agent по его ID
func (r *userAgentRepo) UpdateUserAgentLastAccessedById(ctx context.Context, id int64) (*db.UserAgent, error) {
	return r.queries.UpdateUserAgentLastAccessedById(ctx, id)
}
