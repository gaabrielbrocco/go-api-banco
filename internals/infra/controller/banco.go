package controller

import (
	"encoding/json"
	"database/sql"
	"strconv"
	"net/http"
	"teste/internals/core/domain"
	"teste/internals/core/dto"
	"github.com/go-chi/chi"
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

	input.Nome = strings.TrimSpace(input.Nome)
	if input.Nome == "" {
 	 	http.Error(response, "nome is required", http.StatusBadRequest)
		return
	}

	output, err := controller.bancoUseCase.Create(ctx, &input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func (controller *bancoController) GetByID(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	idString := chi.URLParam(request, "id")
	if idString == "" {
		http.Error(response, "id is required", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	output, err := controller.bancoUseCase.GetByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(response, "record not found", http.StatusNotFound)
			return
		}
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func (controller *bancoController) ListAll(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	
	output, err := controller.bancoUseCase.ListAll(ctx)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	
	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *bancoController) DeleteByID(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	idString := chi.URLParam(request, "id")
	if idString == "" {
		http.Error(response, "id is required", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	output, err := controller.bancoUseCase.DeleteByID(ctx, id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *bancoController) Update(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	idString := chi.URLParam(request, "id")
	if idString == "" {
		http.Error(response, "id is required", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	var input dto.BancoInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	
	input.Nome = strings.TrimSpace(input.Nome)
	if input.Nome == "" {
 	 	http.Error(response, "nome is required", http.StatusBadRequest)
		return
	}

	output, err := controller.bancoUseCase.Update(ctx, id, input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func NewBancoController(bancoUseCase domain.BancoUseCase) domain.BancoController {
	return &bancoController{
		bancoUseCase: bancoUseCase,
	}
}
