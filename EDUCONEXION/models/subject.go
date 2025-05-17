package models

import "gorm.io/gorm"

// Subject representa la tabla de materias
type Subject struct {
    gorm.Model
    Name   string `gorm:"type:varchar(255);not null" json:"name"`
    Code   string `gorm:"type:varchar(50);not null;index:idx_user_code,unique" json:"code"`
    UserID uint   `gorm:"not null;index:idx_user_code,unique" json:"-"`
}
