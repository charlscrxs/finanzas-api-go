package models

type Credenciales struct {
	Email      string `json:"email" binding:"required"`
	Contrasena string `json:"contrasena" binding:"required"`
}
