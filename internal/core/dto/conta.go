package dto

type ContaInput struct {
	UsuarioID int64  `json:"usuario_id" validate:"required,gt=0"`
	BancoID   int64  `json:"banco_id" validate:"required,gt=0"`
	Agencia   string `json:"agencia" validate:"required"`
	Numero    string `json:"numero" validate:"required"`
}
