package repo

import (
	"context"
	"soccer-api/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct{ Pool *pgxpool.Pool }

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := r.Pool.QueryRow(ctx, query, email).Scan(&u.ID, &u.Email, &u.PasswordHash)
	return &u, err
}

func (r *UserRepo) Create(ctx context.Context, email, hash string) (int64, error) {
	var id int64
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	err := r.Pool.QueryRow(ctx, query, email, hash).Scan(&id)
	return id, err
}
