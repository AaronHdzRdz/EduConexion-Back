package routes

import (
    "github.com/gin-gonic/gin"
    "gorm/handlers"
)

func SetupSubjectRoutes(rg *gin.RouterGroup, h *handlers.SubjectHandler) {
    grp := rg.Group("/subjects")
    {
        grp.POST("",    h.Create)   // POST   /api/subjects
        grp.GET("",     h.List)     // GET    /api/subjects
        grp.GET("/:id", h.Get)      // GET    /api/subjects/:id
        grp.PUT("/:id", h.Update)   // PUT    /api/subjects/:id
        grp.DELETE("/:id", h.Delete)// DELETE /api/subjects/:id
    }
}
