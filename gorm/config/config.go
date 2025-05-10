package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config contiene la configuración de la aplicación
type Config struct {
	DBHost     string // Host de Postgres
	DBPort     string // Puerto de Postgres
	DBUser     string // Usuario de Postgres
	DBPassword string // Contraseña de Postgres
	DBName     string // Nombre de la base de datos
	ServerPort string // Puerto en el que arranca el servidor HTTP
}

// Load carga variables de entorno (y .env en desarrollo) y retorna la configuración
func Load() *Config {
	// Carga variables desde archivo .env (si existe) y del entorno
	_ = godotenv.Load()

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "294332"),
		DBName:     getEnv("DB_NAME", "EduConexion"),
		// Usa PORT para el servidor; fallback 3000
		ServerPort: getEnv("PORT", "3000"),
	}
}

// getEnv devuelve el valor de la variable de entorno o el fallback si no existe
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
