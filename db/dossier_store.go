package db

import (
	"context"
	"database/sql"

	"github.com/MoXcz/dossier-org/internal/database"
)

type DossierStore interface {
	GetDossiersFromUserID(context.Context, int64) ([]database.Dossier, error)
	CreateDossier(context.Context, *database.CreateDossierParams) (*database.Dossier, error)
}

type PostgresDossierStore struct {
	db *database.Queries
}

func NewPostgresDossierStore(db *sql.DB) *PostgresDossierStore {
	return &PostgresDossierStore{
		db: database.New(db),
	}
}

func (s *PostgresDossierStore) CreateDossier(ctx context.Context, dossier *database.CreateDossierParams) (*database.Dossier, error) {
	createdDossier, err := s.db.CreateDossier(ctx, *dossier)
	if err != nil {
		return nil, err
	}

	return &createdDossier, nil
}

func (s *PostgresDossierStore) GetDossiersFromUserID(ctx context.Context, id int64) ([]database.Dossier, error) {
	createdDossier, err := s.db.GetDossiersFromUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return createdDossier, nil
}
