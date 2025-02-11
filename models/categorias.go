package models

import "time"

type Categorias struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Nombre      string    `gorm:"type:text" json:"nombre"`
	Descripcion string    `gorm:"type:text" json:"descripcion"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
