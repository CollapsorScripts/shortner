package db

import (
	"context"
	"errors"
	"fmt"
	"shortner/internal/config"
	"shortner/pkg/logger"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const contextTimeout = time.Second * 10

// CreateDatabase - создает базу данных
func CreateDatabase(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable", cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password, cfg.Database.Port)

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	defer conn.Close()

	if cfg.Paths.Migrations.Drop {
		terminateSQL := fmt.Sprintf(`
        SELECT pg_terminate_backend(pid)
        FROM pg_stat_activity
        WHERE datname = '%s' AND pid <> pg_backend_pid();`,
			cfg.Database.Name,
		)

		if _, err := conn.Exec(ctx, terminateSQL); err != nil {
			logger.Warn("Ошибка при завершении активных соединений: %v", err)
		}

		queryDrop := fmt.Sprintf("drop database %s;", cfg.Database.Name)
		if _, err = conn.Exec(ctx, queryDrop); err != nil {
			logger.Warn("Ошибка при попытке удалить базу данных: %v", err)
		}
	}

	query := fmt.Sprintf("create database %s;", cfg.Database.Name)

	_, err = conn.Exec(ctx, query)
	if err != nil {
		if strings.Contains(err.Error(), "already exists (SQLSTATE 42P04)") {
			return nil
		}

		return err
	}

	return nil
}

// Connect - подключение к базе данных
func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	dbCfg := cfg.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbCfg.Host,
		dbCfg.User,
		dbCfg.Password, dbCfg.Name, dbCfg.Port)
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Close - закрывает соединение с базой данных
func Close(conn *pgxpool.Pool) {
	conn.Close()
}

// AutoMigrate - приминение миграции
func AutoMigrate(cfg *config.Config) error {
	path := fmt.Sprintf("file://%s", cfg.Paths.Migrations.Dir)
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
	)

	m, err := migrate.New(path, dsn)
	if err != nil {
		logger.Error("Ошибка при миграции: %v\n", err)
		return err
	}

	if cfg.Paths.Migrations.ReCreate {
		if err := m.Down(); err != nil {
			if err != migrate.ErrNoChange {
				logger.Error("Ошибка при пересоздании таблицы: %v", err)
			}
		}
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("Ошибка при поднятии миграции: %v\n", err)
		return err
	}

	return nil
}

// Drop - удаление таблицы
func Drop(cfg *config.Config) error {
	path := fmt.Sprintf("file://%s", cfg.Paths.Migrations.Dir)
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
	)

	m, err := migrate.New(path, dsn)
	if err != nil {
		logger.Error("ошибка при миграции: %v\n", err)
		return err
	}

	if err := m.Drop(); err != nil {
		if err != migrate.ErrNoChange {
			logger.Error("ошибка при удалении таблицы: %v\n", err)
			return err
		}

		return err
	}

	return nil
}
