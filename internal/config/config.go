package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	LOCAL = "local"
	PROD  = "prod"
)

type Migrations struct {
	Dir      string `yaml:"dir"`
	ReCreate bool   `yaml:"recreate"`
	Drop     bool   `yaml:"drop"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Paths struct {
	LogDir     string      `yaml:"logDir"`
	LogName    string      `yaml:"logName"`
	Migrations *Migrations `yaml:"migrations"`
}

type Config struct {
	Env      string    `yaml:"env" env-default:"local"`
	Paths    *Paths    `yaml:"paths"`
	Database *Database `yaml:"database"`
}

// MustLoad - загружает конфигурацию
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Файл конфигурации по указанному пути отсутствует")
	}

	return MustLoadByPath(path)
}

// MustLoadByPath - загружает конфигурацию по указанному пути
func MustLoadByPath(configPath string) *Config {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("Файл конфигурации не найден: %s", configPath))
	}

	cfg := new(Config)

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		panic(fmt.Sprintf("Ошибка чтения файла конфигурации: %v", err))
	}

	return mustFecthDbEnv(cfg)
}

// mustFecthDbEnv - загружает конфигурацию базы данных из переменных окружения (если окружение - PROD, для локальной конфигурации ничего не изменится)
func mustFecthDbEnv(cfg *Config) *Config {
	if cfg.Env == PROD {
		envArgs := map[string]string{
			"DB_HOST":     "",
			"DB_PORT":     "",
			"DB_USER":     "",
			"DB_PASSWORD": "",
			"DB_NAME":     "",
		}

		for key := range envArgs {
			val := os.Getenv(key)
			if val == "" {
				panic(fmt.Sprintf("%s не указан", key))
			}
			envArgs[key] = val
		}

		port, err := strconv.Atoi(envArgs["DB_PORT"])
		if err != nil {
			panic(fmt.Sprintf("Неверный формат DB_PORT: %v", err))
		}

		cfg.Database = &Database{
			Host:     envArgs["DB_HOST"],
			Port:     port,
			User:     envArgs["DB_USER"],
			Password: envArgs["DB_PASSWORD"],
			Name:     envArgs["DB_NAME"],
		}
	}

	return cfg
}

// fetchConfigPath - парсинг пути к конфигурации из флага или переменной окружения
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		if os.Getenv("CONFIG_PATH") == "" {
			panic("Конфигурация не найдена!")
		}
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
