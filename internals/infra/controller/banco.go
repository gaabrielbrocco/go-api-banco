package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"teste/internals/core/domain"
	"teste/internals/core/dto"
)

type bancoController struct {
	bancoUseCase domain.BancoUseCase
}

func (controller *bancoController) Create(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var input dto.BancoInput

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.bancoUseCase.Create(ctx, &input)
	fmt.Println(output)
	fmt.Println(err)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func NewBancoController(bancoUseCase domain.BancoUseCase) domain.BancoController {
	return &bancoController{
		bancoUseCase: bancoUseCase,
	}
}
