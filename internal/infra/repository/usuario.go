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

func (repository *usuarioRepository) ListAll(ctx context.Context) ([]domain.Usuario, error) {
	query := "SELECT id, nome, email, cpf FROM usuario"
	rows, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []domain.Usuario

	for rows.Next() {
		var usuario domain.Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Cpf); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (repository *usuarioRepository) GetByID(ctx context.Context, id int64) (domain.Usuario, error) {
	var usuario domain.Usuario

	query := "SELECT * FROM usuario WHERE id = $1"

	err := repository.db.QueryRowContext(ctx, query, id).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Cpf)
	if err != nil {
		return domain.Usuario{}, err
	}

	return usuario, nil
}

func NewUsuarioRepository(db *sql.DB) domain.UsuarioRepository {
	return &usuarioRepository{
		db: db,
	}
}
