package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config - основная структура конфига приложения которая агрегирует все другие конфиги
type Config struct {
	App      App
	Postgres Postgres
}

// GetConfig - получает конфиг файла на основе переменных окружения
func GetConfig() (Config, error) {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		return Config{}, fmt.Errorf("read config from env vars: %w", err)
	}
	return conf, nil
}
