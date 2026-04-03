package repository

import (
	"context"
	db "shortner/internal/database/sqlc"
)

type FingerPrintRepository interface {
	CreateFingerPrint(ctx context.Context, statisticsID int64, ip string) (*db.Fingerprint, error)
	GetFingerPrints(ctx context.Context) ([]*db.Fingerprint, error)
	ListFingerPrint(ctx context.Context, statisticsID int64) ([]*db.Fingerprint, error)
	ListFingerPrintByStatisticsId(ctx context.Context, statisticsID int64, limit, offset int32) ([]*db.Fingerprint, error)
	GetFingerPrintByIp(ctx context.Context, ip string) (*db.Fingerprint, error)
	CreateFullFingerPrint(ctx context.Context, statistics_id int64, ip, agent string) (*db.Fingerprint, *db.UserAgent, error)
}
