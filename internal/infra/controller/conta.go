package controller

import (
	"encoding/json"
	"net/http"
	"teste/internal/core/domain"
	"teste/internal/core/dto"
)

type contaController struct {
	contaUseCase domain.ContaUseCase
}

func (controller *contaController) Create(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var input dto.ContaInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.contaUseCase.Create(ctx, &input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func NewContaController(contaUseCase domain.ContaUseCase) domain.ContaController {
	return &contaController{
		contaUseCase: contaUseCase,
	}
}
