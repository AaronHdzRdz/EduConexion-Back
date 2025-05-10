package models

import (
	"gorm.io/gorm"
)

// User representa la tabla de usuarios en la base de datos
type User struct {
	gorm.Model

	// Campos básicos
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`

	// Almacena la contraseña del usuario; se expone en JSON para crear usuarios
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}
