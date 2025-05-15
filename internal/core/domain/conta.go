package domain

import (
	"context"
	"net/http"
	"teste/internal/core/dto"
)

type Conta struct {
	ID        int64  `json:"id" db:"id"`
	UsuarioID int64  `json:"usuario_id" db:"usuario_id"`
	BancoID   int64  `json:"banco_id" db:"banco_id"`
	Agencia   string `json:"agencia" db:"agencia"`
	Numero    string `json:"numero" db:"numero"`
}

type ContaRepository interface {
	Create(context.Context, *dto.ContaInput) (Conta, error)
}

type ContaUseCase interface {
	Create(context.Context, *dto.ContaInput) (Conta, error)
}

type ContaController interface {
	Create(http.ResponseWriter, *http.Request)
}
