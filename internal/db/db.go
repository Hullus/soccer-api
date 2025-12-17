package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBPool(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
}
