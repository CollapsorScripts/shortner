package bootstrap

import (
	"shortner/internal/config"
	"shortner/internal/database/repository"
	"shortner/internal/database/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Services struct {
	URLS        *service.UrlsService
	Statistics  *service.StatisticsService
	FingerPrint *service.FingerPrintService
	UserAgent   *service.UserAgentService
	db          *pgxpool.Pool
}

// InitServices - инициализирует сервисы
func InitServices(cfg *config.Config, dbconn *pgxpool.Pool) *Services {
	// Репозитории
	urlRepo := repository.NewUrlsRepository(dbconn, cfg)
	statRepo := repository.NewStatisticsRepository(dbconn, cfg)
	fpRepo := repository.NewFingerPrintRepository(dbconn, cfg)
	uaRepo := repository.NewUserAgentRepository(dbconn, cfg)

	// Сервисы
	urlService := service.NewUrlsService(urlRepo)
	statService := service.NewStatisticsService(statRepo)
	fpService := service.NewFingerPrintService(fpRepo)
	uaService := service.NewUserAgentService(uaRepo)

	return &Services{
		URLS:        urlService,
		Statistics:  statService,
		FingerPrint: fpService,
		UserAgent:   uaService,
		// -----------
		db: dbconn,
	}
}

// DB - возвращает экземпляр pgxpool.Pool для работы с БД
func (s *Services) DB() *pgxpool.Pool {
	return s.db
}
