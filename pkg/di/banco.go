package di

import (
	"database/sql"
	"teste/internal/core/domain"
	"teste/internal/core/usecase"
	"teste/internal/infra/controller"
	"teste/internal/infra/repository"
)

func NewBancoController(database *sql.DB) domain.BancoController {
	repository := repository.NewBancoRepository(database)
	useCase := usecase.NewBancoUseCase(repository)
	return controller.NewBancoController(useCase)
}
