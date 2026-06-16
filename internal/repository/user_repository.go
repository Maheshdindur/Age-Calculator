package repository

import (
	"context"
	"backend-project/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context) ([]db.User, error)
}

type PostgresUserRepository struct {
	queries *db.Queries
	db      *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &PostgresUserRepository{
		queries: db.New(pool),
		db:      pool,
	}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *PostgresUserRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}
