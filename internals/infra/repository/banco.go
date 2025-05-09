package repository

import (
	"context"
	"database/sql"
	"teste/internals/core/domain"
	"teste/internals/core/dto"
)

type bancoRepository struct {
	db *sql.DB
}

func (repository *bancoRepository) Create(ctx context.Context, input *dto.BancoInput) (domain.Banco, error) {

	var banco domain.Banco

	query := `INSERT INTO banco(nome) VALUES ($1) RETURNING *`

	err := repository.db.QueryRowContext(ctx, query, input.Nome).Scan(&banco.ID, &banco.Nome)
	if err != nil {
		return domain.Banco{}, err
	}

	return banco, nil
}

func (repository *bancoRepository) GetByID(ctx context.Context, id int64) (domain.Banco, error) {

	var banco domain.Banco

	query := "SELECT * FROM banco WHERE id = $1"

	err := repository.db.QueryRowContext(ctx, query, id).Scan(&banco.ID, &banco.Nome)
	if err != nil {
		return domain.Banco{}, err
	}

	return banco, nil
}

func NewBancoRepository(db *sql.DB) domain.BancoRepository {
	return &bancoRepository{
		db: db,
	}
}
