package suite

import (
	"context"
	"fmt"
	"net/http"
	"shortner/internal/bootstrap"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"
	"testing"
	"time"
)

type Suite struct {
	*testing.T                // Потребуется для вызова методов *testing.T внутри Suite
	Cfg        *config.Config // Конфигурация приложения
	BaseURL    string
	Client     *http.Client
	Services   *bootstrap.Services
	HeaderKey  string
	HeaderVal  string
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath(configPath())

	cfg.Paths.Migrations.Dir = "../internal/database/migrations"

	ctx, cancelCtx := context.WithTimeout(context.Background(), 60*time.Minute)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	//Инициализируем базу данных
	if err := db.CreateDatabase(cfg); err != nil {
		panic(fmt.Errorf("ошибка при создании базы данных: %w", err))
	}

	// Применяем миграции
	if err := db.AutoMigrate(cfg); err != nil {
		panic(fmt.Errorf("ошибка при применении миграций: %w", err))
	}

	// Инициализируем соединение с базой данных
	dbConnect, err := db.Connect(cfg)
	if err != nil {
		panic(fmt.Errorf("не удалось установить соединение с базой данных: %w", err))
	}

	//Инициализируем сервисы
	services := bootstrap.InitServices(cfg, dbConnect)

	return ctx, &Suite{
		T:         t,
		Cfg:       cfg,
		BaseURL:   fmt.Sprintf("http://localhost:%d/api/v1/", cfg.Server.Port),
		Client:    &http.Client{Timeout: cfg.Server.Timeout},
		HeaderKey: "Content-Type",
		HeaderVal: "application/json",
		Services:  services,
	}
}

func (s *Suite) SetStandardHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func configPath() string {
	return "../config/local.yaml"
}
