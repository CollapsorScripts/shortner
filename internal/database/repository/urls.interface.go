package repository

import (
	"context"
	db "shortner/internal/database/sqlc"
)

type UrlsRepository interface {
	CreateUrl(ctx context.Context, url, short_url string) (*db.Url, error)
	GetOriginalUrlByShortUrl(ctx context.Context, shortUrl string) (string, error)
	GetOriginalUrlById(ctx context.Context, id int64) (string, error)
	GetUrls(ctx context.Context) ([]*db.Url, error)
	ListUrls(ctx context.Context, limit, offset int32) ([]*db.Url, error)
	DeleteUrlById(ctx context.Context, id int64) error
	DeleteUrlByShortUrl(ctx context.Context, shortUrl string) error
	GetUrlByShortUrl(ctx context.Context, shortUrl string) (*db.Url, error)
}
