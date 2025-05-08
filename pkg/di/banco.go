package di

import (
	"database/sql"
	"teste/internals/core/domain"
	"teste/internals/core/usecase"
	"teste/internals/infra/controller"
	"teste/internals/infra/repository"
)

func NewBancoController(database *sql.DB) domain.BancoController {
	repository := repository.NewBancoRepository(database)
	useCase := usecase.NewBancoUseCase(repository)
	return controller.NewBancoController(useCase)
}
