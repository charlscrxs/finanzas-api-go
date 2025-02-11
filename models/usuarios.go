package models

type Usuarios struct {
	ID         uint   `json:"id"`
	Nombre     string `json:"nombre"`
	Email      string `json:"email"`
	Contrasena string `gorm:"not null" json:"-"`
}
