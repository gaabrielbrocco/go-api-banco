package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teste/internal/core/domain"
	"teste/internal/core/dto"

	"github.com/go-chi/chi"
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

func (controller *contaController) ListByUser(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	idString := chi.URLParam(request, "usuario_id")
	if idString == "" {
		http.Error(response, "usuario_id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.contaUseCase.ListByUser(ctx, id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func NewContaController(contaUseCase domain.ContaUseCase) domain.ContaController {
	return &contaController{
		contaUseCase: contaUseCase,
	}
}
