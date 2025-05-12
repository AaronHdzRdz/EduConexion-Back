package models

import "gorm.io/gorm"

// Student representa la tabla de alumnos
type Student struct {
    gorm.Model
    Name  string `gorm:"type:varchar(255);not null" json:"name"`
    Group string `gorm:"type:varchar(100)"          json:"group"`
    Email string `gorm:"type:varchar(100);unique"   json:"email"`
}
