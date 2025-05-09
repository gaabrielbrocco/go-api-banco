package domain

import (
	"context"
	"net/http"
	"teste/internals/core/dto"
)

type Banco struct {
	ID   int64  `json:"id" db:"id"`
	Nome string `json:"nome" db:"nome"`
}

type BancoRepository interface {
	Create(context.Context, *dto.BancoInput) (Banco, error)
	GetByID(context.Context, int64) (Banco, error)
	ListAll(context.Context) ([]Banco, error)
}

type BancoUseCase interface {
	Create(context.Context, *dto.BancoInput) (Banco, error)
	GetByID(context.Context, int64) (Banco, error)
	ListAll(context.Context) ([]Banco, error)
}

type BancoController interface {
	Create(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
	ListAll(http.ResponseWriter, *http.Request)
}
