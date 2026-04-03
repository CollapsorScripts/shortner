package repository

import (
	"context"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type urlsRepo struct {
	db      *pgxpool.Pool
	cfg     *config.Config
	queries *db.Queries
}

// NewUrlsRepository - создает новый репозиторий для работы с URL
func NewUrlsRepository(conn *pgxpool.Pool, cfg *config.Config) UrlsRepository {
	return &urlsRepo{
		db:      conn,
		cfg:     cfg,
		queries: db.New(conn),
	}
}

// CreateUrl - создает новый URL в базе данных
func (r *urlsRepo) CreateUrl(ctx context.Context, url, short_url string) (*db.Url, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	createdUrl := new(db.Url)

	qtx := r.queries.WithTx(tx)
	for {
		createdUrl, err = qtx.CreateUrl(ctx, db.CreateUrlParams{
			OriginalUrl: url,
			ShortUrl:    short_url,
		})

		if err == nil && createdUrl.ID != 0 {
			break
		}
	}

	_, err = qtx.CreateStatistics(ctx, createdUrl.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return createdUrl, nil
}

// GetOriginalUrlByShortUrl - возвращает оригинальный URL по короткой ссылке
func (r *urlsRepo) GetOriginalUrlByShortUrl(ctx context.Context, shortUrl string) (string, error) {
	return r.queries.GetOriginalUrlByShortUrl(ctx, shortUrl)
}

// GetOriginalUrlById - возвращает оригинальный URL по ID
func (r *urlsRepo) GetOriginalUrlById(ctx context.Context, id int64) (string, error) {
	return r.queries.GetOriginalUrlById(ctx, id)
}

// GetUrls - возвращает все URL из базы данных
func (r *urlsRepo) GetUrls(ctx context.Context) ([]*db.Url, error) {
	return r.queries.GetUrls(ctx)
}

// ListUrls - возвращает все URL из базы данных (пагинация)
func (r *urlsRepo) ListUrls(ctx context.Context, limit, offset int32) ([]*db.Url, error) {
	return r.queries.ListUrls(ctx, db.ListUrlsParams{
		Limit:  limit,
		Offset: offset,
	})
}

// DeleteUrlById - удаляет URL по ID
func (r *urlsRepo) DeleteUrlById(ctx context.Context, id int64) error {
	return r.queries.DeleteUrlById(ctx, id)
}

// DeleteUrlByShortUrl - удаляет URL по короткой ссылке
func (r *urlsRepo) DeleteUrlByShortUrl(ctx context.Context, shortUrl string) error {
	return r.queries.DeleteUrlByShortUrl(ctx, shortUrl)
}
