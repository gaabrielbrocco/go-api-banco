package usecase

import (
	"context"
	"teste/internal/core/domain"
	"teste/internal/core/dto"
)

type contaUseCase struct {
	contaRepository domain.ContaRepository
}

func (usecase *contaUseCase) Create(ctx context.Context, input *dto.ContaInput) (domain.Conta, error) {
	return usecase.contaRepository.Create(ctx, input)
}

func NewContaUseCase(contaRepository domain.ContaRepository) domain.ContaUseCase {
	return &contaUseCase{
		contaRepository: contaRepository,
	}
}
