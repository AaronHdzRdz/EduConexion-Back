package routes

import (
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm/handlers"
)

func SetupUserRoutes(r *gin.Engine, h *handlers.UserHandler) {
    // 1) CORS middleware para aceptar preflights
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // 2) Grupo de rutas CRUD para /users (incluye OPTIONS)
    grp := r.Group("/users")
    {
        grp.OPTIONS("", func(c *gin.Context) {
            c.Status(http.StatusOK)
        })
        grp.POST("", h.Create)
        grp.GET("", h.List)
        grp.GET("/:id", h.Get)
        grp.PUT("/:id", h.Update)
        grp.DELETE("/:id", h.Delete)
    }

    // 3) Rutas de login y su preflight
    r.OPTIONS("/login", func(c *gin.Context) {
        c.Status(http.StatusOK)
    })
    r.POST("/login", h.Login)
}
