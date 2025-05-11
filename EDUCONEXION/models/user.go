package models

import (
	"time"

)

// User representa la tabla de usuarios en la base de datos, sin soft-delete
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Campos básicos
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

// TableName indica a GORM que esta struct usa la tabla "user"
func (User) TableName() string {
	return "user"
}
