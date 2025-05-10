package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database envuelve la conexión GORM
type Database struct {
	DB *gorm.DB
}

// NewDatabase crea y abre la conexión; devuelve error si falla
func NewDatabase() (*Database, error) {
	// Lee todo de ENV, con valores por defecto si quieres
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "294332")
	name := getEnv("DB_NAME", "EduConexion")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error abriendo la base de datos: %w", err)
	}

	return &Database{DB: db}, nil
}

// helper para leer ENV con fallback
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
