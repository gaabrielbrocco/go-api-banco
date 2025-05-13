package repository

import (
	"context"
	"database/sql"
	"errors"
	"teste/internal/core/domain"
	"teste/internal/core/dto"

	"github.com/lib/pq"
)

type usuarioRepository struct {
	db *sql.DB
}

var ErrUsuarioDuplicado = errors.New("email or cpf duplicated")

func (repository *usuarioRepository) Create(ctx context.Context, input *dto.UsuarioInput) (domain.Usuario, error) {
	var usuario domain.Usuario

	query := `INSERT INTO usuario(nome, email, cpf) VALUES ($1, $2, $3) RETURNING *`

	err := repository.db.QueryRowContext(ctx, query, input.Nome, input.Email, input.Cpf).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Cpf)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return domain.Usuario{}, ErrUsuarioDuplicado
		}
		return domain.Usuario{}, err
	}

	return usuario, nil
}

func NewUsuarioRepository(db *sql.DB) domain.UsuarioRepository {
	return &usuarioRepository{
		db: db,
	}
}
