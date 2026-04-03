package repository

import (
	"context"
	db "shortner/internal/database/sqlc"
)

type StatisticsRepository interface {
	CreateStatistics(ctx context.Context, url_id int64) (*db.Statistic, error)
	UpdateLastAccessedById(ctx context.Context, url_id int64) (*db.Statistic, error)
	UpdateLastAccessedByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error)
	IncrementClicksCountByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error)
	IncrementClicksCountById(ctx context.Context, url_id int64) (*db.Statistic, error)
	GetStatistics(ctx context.Context) ([]*db.Statistic, error)
	GetStatisticById(ctx context.Context, url_id int64) (*db.Statistic, error)
	GetStatisticsByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error)
	ListStatistics(ctx context.Context, limit, offset int32) ([]*db.Statistic, error)
}
