package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"shellrean.id/account/internal/database"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	queries := database.New(pool)

	acc, err := queries.GetAccountByID(ctx, "123")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("ACCOUNT-NAME: %s", acc.Name)
}
