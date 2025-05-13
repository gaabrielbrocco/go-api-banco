package usecase

import (
	"context"
	"teste/internal/core/domain"
	"teste/internal/core/dto"
)

type usuarioUseCase struct {
	usuarioRepository domain.UsuarioRepository
}

func (usecase *usuarioUseCase) Create(ctx context.Context, input *dto.UsuarioInput) (domain.Usuario, error) {
	return usecase.usuarioRepository.Create(ctx, input)
}

func (usecase *usuarioUseCase) ListAll(ctx context.Context) ([]domain.Usuario, error) {
	return usecase.usuarioRepository.ListAll(ctx)
}

func (usecase *usuarioUseCase) GetByID(ctx context.Context, id int64) (domain.Usuario, error) {
	return usecase.usuarioRepository.GetByID(ctx, id)
}

func NewUsuarioUseCase(usuarioRepository domain.UsuarioRepository) domain.UsuarioUseCase {
	return &usuarioUseCase{
		usuarioRepository: usuarioRepository,
	}
}
