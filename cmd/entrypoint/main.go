package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shortner/internal/bootstrap"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"
	"shortner/internal/server"
	"shortner/pkg/logger"
	"shortner/utils"
	"time"
)

func main() {
	//Инициализация конфигурации
	cfg := config.MustLoad()

	//Инициализируем логгер
	if err := logger.New(cfg); err != nil {
		fmt.Printf("Ошибка при инициализации логера: %v\n", err)
	}

	logger.Info("Загруженная конфигурация: \n%s", utils.ToJSON(cfg))

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

	services := bootstrap.InitServices(cfg, dbConnect)

	//Создаем контекст
	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	{
		//Инициализируем сервер
		srv := server.NewRouter(cfg, services)

		// Запуск сервера в отдельном потоке
		go func() {
			port := fmt.Sprintf(":%d", cfg.Server.Port)

			logger.Info("Сервер запущен на адресе: %s", port)
			if err := srv.Listen(port); err != nil {
				logger.Warn("Ошибка запуска сервера: %v", err)
			}
		}()

		<-rootCtx.Done()
		logger.Warn("Получен SIGINT, выключаемся...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// / Ожидаем завершения сервера
		if err := srv.ShutdownWithContext(shutdownCtx); err != nil {
			logger.Error("Ошибка при выключении сервера api: %v", err)
		}

		dbConnect.Close()
	}
}
