package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shortner/internal/config"
	db "shortner/internal/database/sqlc"
	"shortner/pkg/logger"
	"shortner/utils"
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

	// TODO: потом убрать
	_ = dbConnect

	//Создаем контекст
	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	{
		<-rootCtx.Done()
		logger.Warn("Получен SIGINT, выключаемся...")
	}
}
