package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config - основная структура конфига приложения которая агрегирует все другие конфиги
type Config struct {
	App      App
	Postgres Postgres
}

func GetConfig() (Config, error) {
	var conf Config

	// Загрузка .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env файл или он отсутствует")
	}

	if err := envconfig.Process("", &conf); err != nil {
		return Config{}, fmt.Errorf("read config from env vars: %w", err)
	}

	return conf, nil
}
