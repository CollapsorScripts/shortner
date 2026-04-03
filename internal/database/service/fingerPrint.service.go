package service

import (
	"context"
	"shortner/internal/database/repository"
	db "shortner/internal/database/sqlc"
)

type FingerPrintService struct {
	repo repository.FingerPrintRepository
}

// NewFingerPrintService - создает новый сервис для работы с fingerprint.
func NewFingerPrintService(repo repository.FingerPrintRepository) *FingerPrintService {
	return &FingerPrintService{repo: repo}
}

// GetFingerPrints - возвращает все fingerprint
func (s *FingerPrintService) GetFingerPrints(ctx context.Context) ([]*db.Fingerprint, error) {
	return s.repo.GetFingerPrints(ctx)
}

// ListFingerPrint - возвращает все fingerprint для пользователя
func (s *FingerPrintService) ListFingerPrint(ctx context.Context, statisticsID int64) ([]*db.Fingerprint, error) {
	return s.repo.GetFingerPrints(ctx)
}

// ListFingerPrintByStatisticsId - возвращает все fingerprint для пользователя по statisticsID
func (s *FingerPrintService) ListFingerPrintByStatisticsId(ctx context.Context, statisticsID int64, limit, offset int32) ([]*db.Fingerprint, error) {
	return s.repo.ListFingerPrintByStatisticsId(ctx, statisticsID, limit, offset)
}

// GetFingerPrintByIp - возвращает fingerprint по ip
func (s *FingerPrintService) GetFingerPrintByIp(ctx context.Context, ip string) (*db.Fingerprint, error) {
	return s.repo.GetFingerPrintByIp(ctx, ip)
}
