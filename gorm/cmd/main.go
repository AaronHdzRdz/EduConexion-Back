package main

import (
    "fmt"
    "log"

    "github.com/gin-gonic/gin"

    "gorm/config"
    "gorm/database"
    "gorm/handlers"
    "gorm/routes"
    "gorm/services"
    "gorm/models"
)

func main() {
    cfg := config.Load()

    dbConn, err := database.NewDatabase()
    if err != nil {
        log.Fatalf("❌ Error DB: %v", err)
    }
    fmt.Println("✅ Conectado a Postgres")

    if err := dbConn.DB.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("❌ Migración fallida: %v", err)
    }

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    userSvc := services.NewUserService(dbConn.DB)
    userH := handlers.NewUserHandler(userSvc)
    routes.SetupUserRoutes(r, userH)

    log.Printf("Servidor escuchando en :%s", cfg.ServerPort)
    log.Fatal(r.Run(":" + cfg.ServerPort))
}
