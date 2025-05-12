package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm/config"
	"gorm/database"
	"gorm/handlers"
	"gorm/models"
	"gorm/routes"
	"gorm/services"
)

func main() {
	// 1) Carga .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env no encontrado, usando vars de entorno")
	}

	// 2) Carga configuración (solo ServerPort)
	cfg := config.Load()

	// 3) Conecta a la base de datos
	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("❌ Error DB: %v", err)
	}
	fmt.Println("✅ Conectado a Postgres")

	// 4) Auto-migraciones de todos los modelos
	if err := dbConn.DB.AutoMigrate(
		&models.User{},
		&models.Student{},
		//&models.Subject{},
		//&models.Grade{},
	); err != nil {
		log.Fatalf("❌ Migración fallida: %v", err)
	}

	// 5) Crea el router Gin
	r := gin.Default()

	// 6) Ruta de prueba
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 7) Inicializa servicios y handlers

	// Usuarios
	userSvc := services.NewUserService(dbConn.DB)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET no configurado")
	}
	userH := handlers.NewUserHandler(userSvc, jwtSecret)
	routes.SetupUserRoutes(r, userH)

	// Estudiantes
	studentSvc := services.NewStudentService(dbConn.DB)
	studentH := handlers.NewStudentHandler(studentSvc)
	routes.SetupStudentRoutes(r, studentH)

	// Materias
	// subjectSvc := services.NewSubjectService(dbConn.DB)
	// subjectH := handlers.NewSubjectHandler(subjectSvc)
	// routes.SetupSubjectRoutes(r, subjectH)

	// Calificaciones
	// gradeSvc := services.NewGradeService(dbConn.DB)
	// gradeH := handlers.NewGradeHandler(gradeSvc)
	// routes.SetupGradeRoutes(r, gradeH)

	// 8) Arranca el servidor en todas las interfaces
	addr := "0.0.0.0:" + cfg.ServerPort
	log.Printf("Servidor escuchando en %s (todas las interfaces)", addr)
	log.Fatal(r.Run(addr))
}
