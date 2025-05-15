package repository

import (
	"context"
	"database/sql"
	"teste/internal/core/domain"
	"teste/internal/core/dto"
)

type contaRepository struct {
	db *sql.DB
}

func (repository *contaRepository) Create(ctx context.Context, input *dto.ContaInput) (domain.Conta, error) {
	var conta domain.Conta

	query := "INSERT INTO conta(usuario_id, banco_id, agencia, numero) VALUES ($1, $2, $3, $4) RETURNING *"

	err := repository.db.QueryRowContext(ctx, query, input.UsuarioID, input.BancoID, input.Agencia, input.Numero).Scan(&conta.ID, &conta.UsuarioID, &conta.BancoID, &conta.Agencia, &conta.Numero)
	if err != nil {
		return domain.Conta{}, err
	}

	return conta, nil
}

func NewContaRepository(db *sql.DB) domain.ContaRepository {
	return &contaRepository{
		db: db,
	}
}
