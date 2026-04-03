package service

import (
	"context"
	"shortner/internal/config"
	"shortner/internal/database/repository"
	db "shortner/internal/database/sqlc"
	"shortner/pkg/shortener"
)

type UrlsService struct {
	repo repository.UrlsRepository
}

// NewUrlsService - создает новый сервис для работы с URL.
func NewUrlsService(repo repository.UrlsRepository) *UrlsService {
	return &UrlsService{repo: repo}
}

// CreateUrl - создает новый URL в базе данных
func (s *UrlsService) CreateUrl(ctx context.Context, cfg *config.Config, url string) (*db.Url, error) {
	shortUrl, err := shortener.GenerateShortUrl(cfg, url)
	if err != nil {
		return nil, err
	}

	return s.repo.CreateUrl(ctx, url, shortUrl)
}

// GetOriginalUrlByShortUrl - возвращает оригинальный URL по короткой ссылке
func (s *UrlsService) GetOriginalUrlByShortUrl(ctx context.Context, shortUrl string) (string, error) {
	return s.repo.GetOriginalUrlByShortUrl(ctx, shortUrl)
}

// GetOriginalUrlById - возвращает оригинальный URL по ID
func (s *UrlsService) GetOriginalUrlById(ctx context.Context, id int64) (string, error) {
	return s.repo.GetOriginalUrlById(ctx, id)
}

// GetUrls - возвращает все URL из базы данных
func (s *UrlsService) GetUrls(ctx context.Context) ([]*db.Url, error) {
	return s.repo.GetUrls(ctx)
}

// ListUrls - возвращает все URL из базы данных (пагинация)
func (s *UrlsService) ListUrls(ctx context.Context, limit, offset int32) ([]*db.Url, error) {
	return s.repo.ListUrls(ctx, limit, offset)
}

// DeleteUrlById - удаляет URL по ID
func (s *UrlsService) DeleteUrlById(ctx context.Context, id int64) error {
	return s.repo.DeleteUrlById(ctx, id)
}

// DeleteUrlByShortUrl - удаляет URL по короткой ссылке
func (s *UrlsService) DeleteUrlByShortUrl(ctx context.Context, shortUrl string) error {
	return s.repo.DeleteUrlByShortUrl(ctx, shortUrl)
}
