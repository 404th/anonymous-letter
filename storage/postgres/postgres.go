package postgres

import (
	"context"
	"fmt"

	"github.com/404th/anonymous-letter/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(ctx context.Context, cfg config.Config, lg *logrus.Logger) (postgres *Postgres, err error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	config.ConnConfig.LogLevel = pgx.LogLevelInfo

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		db: pool,
	}, err
}

func (ps *Postgres) CloseDB() {
	ps.db.Close()
}
