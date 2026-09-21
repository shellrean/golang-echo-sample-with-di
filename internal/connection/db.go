package connection

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
	"shellrean.id/account/internal/config"
	"shellrean.id/account/internal/database"
)

func NewDatabase(i do.Injector) (*database.Queries, error) {
	cnf := do.MustInvoke[*config.Config](i)

	pool, err := pgxpool.New(context.Background(), cnf.Database.Dsn)
	if err != nil {
		return nil, err
	}
	return database.New(pool), nil
}
