package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"teste/internal/core/domain"
	"teste/internal/core/dto"
	"teste/internal/infra/repository"

	"github.com/go-chi/chi"
)

type usuarioController struct {
	usuarioUseCase domain.UsuarioUseCase
}

func (controller *usuarioController) Create(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var input dto.UsuarioInput
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	input.Nome = strings.TrimSpace(input.Nome)
	input.Email = strings.TrimSpace(input.Email)
	input.Cpf = strings.TrimSpace(input.Cpf)
	if input.Nome == "" || input.Email == "" || input.Cpf == "" {
		http.Error(response, "nome, email and cpf is required", http.StatusBadRequest)
		return
	}

	output, err := controller.usuarioUseCase.Create(ctx, &input)
	if err != nil {
		if errors.Is(err, repository.ErrUsuarioDuplicado) {
			http.Error(response, "email or cpf already registered", http.StatusConflict)
			return
		}
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func (controller *usuarioController) ListAll(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	output, err := controller.usuarioUseCase.ListAll(ctx)
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

func (controller *usuarioController) GetByID(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	idString := chi.URLParam(request, "id")
	if idString == "" {
		http.Error(response, "id is required", http.StatusBadRequest)
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	output, err := controller.usuarioUseCase.GetByID(ctx, id)
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

func NewUsuarioController(usuarioUseCase domain.UsuarioUseCase) domain.UsuarioController {
	return &usuarioController{
		usuarioUseCase: usuarioUseCase,
	}
}
