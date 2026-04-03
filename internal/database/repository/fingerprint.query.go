package repository

import (
	"context"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type fingerPrintRepo struct {
	db      *pgxpool.Pool
	cfg     *config.Config
	queries *db.Queries
}

// NewFingerPrintRepository - создает новый репозиторий для работы с fingerprint
func NewFingerPrintRepository(conn *pgxpool.Pool, cfg *config.Config) FingerPrintRepository {
	return &fingerPrintRepo{
		db:      conn,
		cfg:     cfg,
		queries: db.New(conn),
	}
}

// CreateFingerPrint - создает новый fingerprint для пользователя
func (r *fingerPrintRepo) CreateFingerPrint(ctx context.Context, statisticsID int64, ip string) (*db.Fingerprint, error) {
	return r.queries.CreateFingerPrint(ctx, db.CreateFingerPrintParams{
		StatisticsID: statisticsID,
		Ip:           ip,
	})
}

// GetFingerPrints - возвращает все fingerprint
func (r *fingerPrintRepo) GetFingerPrints(ctx context.Context) ([]*db.Fingerprint, error) {
	return r.queries.GetFingerPrints(ctx)
}

// ListFingerPrint - возвращает все fingerprint для пользователя
func (r *fingerPrintRepo) ListFingerPrint(ctx context.Context, statisticsID int64) ([]*db.Fingerprint, error) {
	return r.queries.GetFingerPrints(ctx)
}

// ListFingerPrintByStatisticsId - возвращает все fingerprint для пользователя по statisticsID
func (r *fingerPrintRepo) ListFingerPrintByStatisticsId(ctx context.Context, statisticsID int64, limit, offset int32) ([]*db.Fingerprint, error) {
	return r.queries.ListFingerPrintByStatisticsId(ctx, db.ListFingerPrintByStatisticsIdParams{
		StatisticsID: statisticsID,
		Limit:        limit,
		Offset:       offset,
	})
}

// GetFingerPrintByIp - возвращает fingerprint по ip
func (r *fingerPrintRepo) GetFingerPrintByIp(ctx context.Context, ip string) (*db.Fingerprint, error) {
	return r.queries.GetFingerPrintByIp(ctx, ip)
}
