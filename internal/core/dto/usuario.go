package dto

type UsuarioInput struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}
