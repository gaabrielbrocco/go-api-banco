package usecase

import (
	"context"
	"teste/internal/core/domain"
	"teste/internal/core/dto"
)

type bancoUseCase struct {
	bancoRepository domain.BancoRepository
}

func (usecase *bancoUseCase) Create(ctx context.Context, input *dto.BancoInput) (domain.Banco, error) {
	return usecase.bancoRepository.Create(ctx, input)
}

func (usecase *bancoUseCase) GetByID(ctx context.Context, id int64) (domain.Banco, error) {
	return usecase.bancoRepository.GetByID(ctx, id)
}

func (usecase *bancoUseCase) ListAll(ctx context.Context) ([]domain.Banco, error) {
	return usecase.bancoRepository.ListAll(ctx)
}

func (usecase *bancoUseCase) DeleteByID(ctx context.Context, id int64) (domain.Banco, error) {
	return usecase.bancoRepository.DeleteByID(ctx, id)
}

func (usecase *bancoUseCase) Update(ctx context.Context, id int64, input dto.BancoInput) (domain.Banco, error) {
	return usecase.bancoRepository.Update(ctx, id, input)
}

func NewBancoUseCase(bancoRepository domain.BancoRepository) domain.BancoUseCase {
	return &bancoUseCase{
		bancoRepository: bancoRepository,
	}
}
