package models

type Categorias struct {
	ID          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	UsuarioID   uint   `json:"usuario_id"`
}
