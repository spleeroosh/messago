package bootstrap

import (
	"context"
	"fmt"
	"github.com/spleeroosh/messago/internal/config"
	migrator "github.com/spleeroosh/messago/internal/pkg"
	"github.com/spleeroosh/messago/migrations"
	"strings"

	"github.com/jackc/pgx/v5/stdlib"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func newPostgresClient(lc fx.Lifecycle, conf config.Config) (*pgxpool.Pool, error) {
	fmt.Printf("POSTGRE START %v\n", conf)
	goqu.SetDefaultPrepared(true)
	connString := strings.Join([]string{
		fmt.Sprintf("user=%s", conf.Postgres.User),
		fmt.Sprintf("password=%s", conf.Postgres.Password),
		fmt.Sprintf("dbname=%s", conf.Postgres.Database),
		fmt.Sprintf("host=%s", conf.Postgres.Host),
		fmt.Sprintf("port=%d", conf.Postgres.Port),
		fmt.Sprintf("sslmode=%s", conf.Postgres.SSLMode),
		fmt.Sprintf("connect_timeout=%d", conf.Postgres.ConnTimeout),
		fmt.Sprintf("pool_max_conns=%d", conf.Postgres.MaxConn),
		fmt.Sprintf("pool_max_conn_lifetime=%s", conf.Postgres.MaxConnLifeTime),
		fmt.Sprintf("pool_max_conn_idle_time=%s", conf.Postgres.MaxConnIdleTime),
	}, " ")

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("create pgxpool: %w", err)
	}
	fmt.Println("POOL CREATED")
	lc.Append(fx.StopHook(func() {
		pool.Close()
	}))
	fmt.Println("POOL PRE MIGRATE")
	err = migrate(pool)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("POOL MIGRATEd")

	return pool, nil
}

func migrate(pool *pgxpool.Pool) error {
	fmt.Println("MIGRATES ARE STARTED")
	return migrator.NewMigrator(stdlib.OpenDBFromPool(pool), migrations.FS).Up()
}
