package models

import "time"

type Movimiento struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Tipo        string    `gorm:"type:varchar(20);not null" json:"tipo"` // 'ingreso' o 'gasto'
	Monto       float64   `gorm:"not null" json:"monto"`
	Descripcion string    `gorm:"type:text" json:"descripcion,omitempty"` // Descripción opcional
	Fecha       time.Time `gorm:"not null;default:current_timestamp" json:"fecha"`

	UsuarioID uint     `gorm:"not null" json:"usuario_id"`
	Usuario   Usuarios `gorm:"foreignKey:UsuarioID" json:"-"`

	CategoriaID *uint       `json:"categoria_id,omitempty"`
	Categoria   *Categorias `gorm:"foreignKey:CategoriaID" json:"categoria,omitempty"`
}
