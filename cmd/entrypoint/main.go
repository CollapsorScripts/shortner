package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shortner/internal/config"
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

	//Создаем контекст
	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	{
		<-rootCtx.Done()
		logger.Warn("Получен SIGINT, выключаемся...")
	}
}
