package routes

import (
    "github.com/gin-gonic/gin"
    "gorm/handlers"
)

func SetupUserRoutes(r *gin.Engine, h *handlers.UserHandler) {
    // Grupo de rutas para CRUD de usuarios
    grp := r.Group("/users")
    {
        grp.POST("", h.Create)        // Crear usuario
        grp.GET("", h.List)           // Listar todos
        grp.GET("/:id", h.Get)        // Obtener por ID
        grp.PUT("/:id", h.Update)     // Actualizar usuario
        grp.DELETE("/:id", h.Delete)  // Eliminar usuario
    }

    // Ruta de login (autenticación)
    r.POST("/login", h.Login)
}
