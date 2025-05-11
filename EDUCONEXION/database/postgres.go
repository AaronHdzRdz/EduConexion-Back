package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// 1) Tables singulares: usa "user" en vez de "users", etc.
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// 2) Habilita logging para ver las queries en consola
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("error abriendo la base de datos: %w", err)
	}

	// 3) Si quieres migrar aquí mismo tus modelos (opcional)
	// import "tu_modulo/models"
	// if err := db.AutoMigrate(&models.User{}, &models.Materia{}, &models.Alumno{}); err != nil {
	//     return nil, fmt.Errorf("error en migraciones: %w", err)
	// }

	return &Database{DB: db}, nil
}

// helper para leer ENV con fallback
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
