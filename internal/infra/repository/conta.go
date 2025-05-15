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

func (repository *contaRepository) ListByUser(ctx context.Context, id int64) ([]domain.Conta, error) {
	query := `
	SELECT c.id, c.usuario_id, c.banco_id, b.nome, c.agencia, c.numero
	FROM conta c
	JOIN banco b ON c.banco_id = b.id
	WHERE c.usuario_id = $1
	`

	rows, err := repository.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contas []domain.Conta

	for rows.Next() {
		var conta domain.Conta
		if err := rows.Scan(&conta.ID, &conta.UsuarioID, &conta.BancoID, &conta.NomeBanco, &conta.Agencia, &conta.Numero); err != nil {
			return nil, err
		}

		contas = append(contas, conta)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contas, nil
}

func NewContaRepository(db *sql.DB) domain.ContaRepository {
	return &contaRepository{
		db: db,
	}
}
