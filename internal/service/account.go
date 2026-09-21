package service

import (
	"context"

	"github.com/samber/do/v2"
	"shellrean.id/account/internal/database"
)

type Account struct {
	db *database.Queries
}

func NewAccount(i do.Injector) (Account, error) {
	return Account{
		db: do.MustInvoke[*database.Queries](i),
	}, nil
}

func (a Account) Show(ctx context.Context, id string) (database.Account, error) {
	return a.db.GetAccountByID(ctx, id)
}
