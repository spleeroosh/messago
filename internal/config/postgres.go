package config

import (
	"time"
)

type Postgres struct {
	Host            string        `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port            int           `envconfig:"POSTGRES_PORT" default:"5432"`
	User            string        `envconfig:"POSTGRES_USER" default:"root"`
	Password        string        `envconfig:"POSTGRES_PASSWORD" default:"root"`
	Database        string        `envconfig:"POSTGRES_DATABASE" default:"messago-db"`
	SSLMode         string        `envconfig:"POSTGRES_SSLMODE" default:"verify-full"`
	ConnTimeout     int           `envconfig:"POSTGRES_CONNTIMEOUT" default:"5"`
	MaxConn         int           `envconfig:"POSTGRES_MAXCONN" default:"100"`
	MaxConnLifeTime time.Duration `envconfig:"POSTGRES_MAXCONN_LIFETIME" default:"25m"`
	MaxConnIdleTime time.Duration `envconfig:"POSTGRES_MAXCONN_IDLETIME" default:"5m"`
}
