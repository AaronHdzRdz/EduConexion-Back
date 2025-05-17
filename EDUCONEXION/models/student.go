// models/student.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(255);not null"   json:"name"`
	Group  string `gorm:"type:varchar(100)"            json:"group"`
	Email  string `gorm:"type:varchar(100);not null;index:idx_user_email,unique" json:"email"`
	UserID uint   `gorm:"not null;index:idx_user_email,unique" json:"user_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
