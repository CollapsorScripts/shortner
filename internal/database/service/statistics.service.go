package service

import (
	"context"
	"shortner/internal/database/repository"
	db "shortner/internal/database/sqlc"
)

type StatisticsService struct {
	repo repository.StatisticsRepository
}

// NewStatisticsService - создает новый сервис для работы с статистикой.
func NewStatisticsService(repo repository.StatisticsRepository) *StatisticsService {
	return &StatisticsService{repo: repo}
}

// CreateStatistics - создает новую запись статистики.
func (s *StatisticsService) CreateStatistics(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.CreateStatistics(ctx, url_id)
}

// UpdateLastAccessedById - обновляет поле last_accessed для записи статистики по url_id.
func (s *StatisticsService) UpdateLastAccessedById(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.UpdateLastAccessedById(ctx, url_id)
}

// UpdateLastAccessedByUrlId - обновляет поле last_accessed для записи статистики по url_id.
func (s *StatisticsService) UpdateLastAccessedByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.UpdateLastAccessedByUrlId(ctx, url_id)
}

// IncrementClicksCountByUrlId - увеличивает счетчик кликов для записи статистики по url_id.
func (s *StatisticsService) IncrementClicksCountByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.IncrementClicksCountByUrlId(ctx, url_id)
}

// IncrementClicksCountById - увеличивает счетчик кликов для записи статистики по url_id.
func (s *StatisticsService) IncrementClicksCountById(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.IncrementClicksCountById(ctx, url_id)
}

// GetStatistics - возвращает статистику по url_id.
func (s *StatisticsService) GetStatistics(ctx context.Context) ([]*db.Statistic, error) {
	return s.repo.GetStatistics(ctx)
}

// GetStatisticById - возвращает статистику по url_id.
func (s *StatisticsService) GetStatisticById(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.GetStatisticById(ctx, url_id)
}

// GetStatisticsByUrlId - возвращает статистику по url_id.
func (s *StatisticsService) GetStatisticsByUrlId(ctx context.Context, url_id int64) (*db.Statistic, error) {
	return s.repo.GetStatisticsByUrlId(ctx, url_id)
}

// ListStatistics - возвращает список статистики (пагинация)
func (s *StatisticsService) ListStatistics(ctx context.Context, limit, offset int32) ([]*db.Statistic, error) {
	return s.repo.ListStatistics(ctx, limit, offset)
}
