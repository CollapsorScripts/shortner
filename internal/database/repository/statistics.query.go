package repository

import (
	"context"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type statisticsRepo struct {
	db      *pgxpool.Pool
	cfg     *config.Config
	queries *db.Queries
}

// NewStatisticsRepository - создает новый репозиторий статистики.
func NewStatisticsRepository(conn *pgxpool.Pool, cfg *config.Config) StatisticsRepository {
	return &statisticsRepo{
		db:      conn,
		cfg:     cfg,
		queries: db.New(conn),
	}
}

// CreateStatistics - создает новую запись статистики.
func (r *statisticsRepo) CreateStatistics(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.CreateStatistics(ctx, url_id)
}

// UpdateLastAccessedById - обновляет поле last_accessed для записи статистики по url_id.
func (r *statisticsRepo) UpdateLastAccessedById(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.UpdateLastAccessedById(ctx, url_id)
}

// UpdateLastAccessedByUrlId - обновляет поле last_accessed для записи статистики по url_id.
func (r *statisticsRepo) UpdateLastAccessedByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.UpdateLastAccessedByUrlId(ctx, url_id)
}

// IncrementClicksCountByUrlId - увеличивает счетчик кликов для записи статистики по url_id.
func (r *statisticsRepo) IncrementClicksCountByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.IncrementClicksCountByUrlId(ctx, url_id)
}

// IncrementClicksCountById - увеличивает счетчик кликов для записи статистики по url_id.
func (r *statisticsRepo) IncrementClicksCountById(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.IncrementClicksCountById(ctx, url_id)
}

// GetStatistics - возвращает статистику по url_id.
func (r *statisticsRepo) GetStatistics(ctx context.Context) ([]*db.Statistic, error) {
	return r.queries.GetStatistics(ctx)
}

// GetStatisticById - возвращает статистику по url_id.
func (r *statisticsRepo) GetStatisticById(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.GetStatisticById(ctx, url_id)
}

// GetStatisticsByUrlId - возвращает статистику по url_id.
func (r *statisticsRepo) GetStatisticsByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return r.queries.GetStatisticsByUrlId(ctx, url_id)
}

// ListStatistics - возвращает список статистики (пагинация)
func (r *statisticsRepo) ListStatistics(ctx context.Context, limit, offset int32) ([]*db.Statistic, error) {
	return r.queries.ListStatistics(ctx, db.ListStatisticsParams{
		Limit:  limit,
		Offset: offset,
	})
}
