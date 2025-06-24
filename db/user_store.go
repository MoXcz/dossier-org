package db

import (
	"context"
	"database/sql"

	"github.com/MoXcz/dossier-org/internal/database"
)

type UserStore interface {
	GetUserByID(context.Context, int64) (*database.User, error)
	GetUsers(context.Context) ([]database.User, error)
	CreateUser(context.Context, *database.CreateUserParams) (*database.User, error)
	Drop(context.Context) error
}

type PostgresUserStore struct {
	db *database.Queries
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: database.New(db),
	}
}

func (s *PostgresUserStore) CreateUser(ctx context.Context, user *database.CreateUserParams) (*database.User, error) {
	createdUser, err := s.db.CreateUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (s *PostgresUserStore) GetUserByID(ctx context.Context, id int64) (*database.User, error) {
	createdUser, err := s.db.GetUserFromID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (s *PostgresUserStore) GetUsers(ctx context.Context) ([]database.User, error) {
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *PostgresUserStore) Drop(ctx context.Context) error {
	return s.db.DeleteUsers(ctx)
}
