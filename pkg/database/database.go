package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Initializes database connection
// Panic if connection can't be established
func ConnectDB() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return dbpool
}
