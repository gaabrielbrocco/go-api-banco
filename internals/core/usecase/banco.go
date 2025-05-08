package usecase

import (
	"context"
	"teste/internals/core/domain"
	"teste/internals/core/dto"
)

type bancoUseCase struct {
	bancoRepository domain.BancoRepository
}

func (usecase *bancoUseCase) Create(ctx context.Context, input *dto.BancoInput) (*domain.Banco, error) {
	return usecase.bancoRepository.Create(ctx, input)
}

func NewBancoUseCase(bancoRepository domain.BancoRepository) domain.BancoUseCase {
	return &bancoUseCase{
		bancoRepository: bancoRepository,
	}
}
