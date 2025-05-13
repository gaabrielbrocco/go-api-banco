package di

import (
	"database/sql"
	"teste/internal/core/domain"
	"teste/internal/core/usecase"
	"teste/internal/infra/controller"
	"teste/internal/infra/repository"
)

func NewUsuarioController(database *sql.DB) domain.UsuarioController {
	repository := repository.NewUsuarioRepository(database)
	useCase := usecase.NewUsuarioUseCase(repository)
	return controller.NewUsuarioController(useCase)
}
