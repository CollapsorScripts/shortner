package repository

import (
	"context"
	db "shortner/internal/database/sqlc"
)

type UserAgentRepository interface {
	CreateUserAgent(ctx context.Context, fingerprintID int64, userAgent string) (*db.UserAgent, error)
	GetUserAgents(ctx context.Context) ([]*db.UserAgent, error)
	GetUserAgentById(ctx context.Context, id int64) (*db.UserAgent, error)
	GetUserAgentsByFingerprintId(ctx context.Context, fingerprintID int64) ([]*db.UserAgent, error)
	UpdateUserAgentLastAccessedById(ctx context.Context, id int64) (*db.UserAgent, error)
}
