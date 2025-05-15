package di

import (
	"database/sql"
	"teste/internal/core/domain"
	"teste/internal/core/usecase"
	"teste/internal/infra/controller"
	"teste/internal/infra/repository"
)

func NewContaController(database *sql.DB) domain.ContaController {
	repository := repository.NewContaRepository(database)
	useCase := usecase.NewContaUseCase(repository)
	return controller.NewContaController(useCase)
}
