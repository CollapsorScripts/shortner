.PHONY: all

# Точка входа
ENTRY_POINT := "./cmd/entrypoint"

# Название бинарного файла и папка куда он будет скомпилирован
BIN_NAME := "shortner"
BIN_DIR := "bin"

# Пути к конфигурациям
CFG_LOCAL := "--config=./config/local.yaml"
CFG_PROD := "--config=./config/prod.yaml"
#

# Аргументы компиляции
LDFLAGS:=-ldflags="-s -w"
CGO=0
#

# Docker
CMD_UP_DOCKER := "docker compose up -d"
CMD_RM_DOCKER := "docker compose stop && docker compose rm -f"
#

all: build run_local

build:
	@echo "Компиляция $(BIN_NAME)"
	@CGO_ENABLED=$(CGO) go build $(LDFLAGS) -o $(BIN_DIR)/$(BIN_NAME) $(ENTRY_POINT)
	@echo "Скомпилировано в папку: $(BIN_DIR)"

run_local:
	@echo "Запуск локального окружения: $(BIN_DIR)/$(BIN_NAME) $(CFG_LOCAL)"
	@$(BIN_DIR)/$(BIN_NAME) $(CFG_LOCAL)

run_prod:
	@echo "Запуск продакшен окружения: $(BIN_DIR)/$(BIN_NAME) $(CFG_PROD)"
	@$(BIN_DIR)/$(BIN_NAME) $(CFG_PROD)

gen_sqlc:
	@echo "Генерация sqlc..."
	@sqlc generate
	@echo "Генерация sqlc завершена!"

update:
	@echo "Обновление пакетов..."
	@go get -u ./...
	@go mod tidy
	@echo "Обновление завершено!"
