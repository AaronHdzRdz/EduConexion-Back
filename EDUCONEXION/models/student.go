// models/student.go
package models

import "gorm.io/gorm"

// Student representa la tabla de alumnos y a qué usuario pertenecen
type Student struct {
    gorm.Model
    Name   string `gorm:"type:varchar(255);not null" json:"name"`
    Group  string `gorm:"type:varchar(100)"          json:"group"`
    Email  string `gorm:"type:varchar(100);not null;index:idx_user_email,unique" json:"email"`
    UserID uint   `gorm:"not null;index:idx_user_email,unique" json:"-"`
}
