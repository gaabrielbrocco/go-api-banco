package domain

import (
	"context"
	"net/http"
	"teste/internal/core/dto"
)

type Usuario struct {
	ID    int64  `json:"id" db:"id"`
	Nome  string `json:"nome" db:"nome"`
	Email string `json:"email" db:"email"`
	Cpf   string `json:"cpf" db:"cpf"`
}

type UsuarioRepository interface {
	Create(context.Context, *dto.UsuarioInput) (Usuario, error)
}

type UsuarioUseCase interface {
	Create(context.Context, *dto.UsuarioInput) (Usuario, error)
}

type UsuarioController interface {
	Create(http.ResponseWriter, *http.Request)
}
