package routes

import (
    "github.com/gin-gonic/gin"
    "gorm/handlers"
)

func SetupStudentRoutes(r *gin.Engine, h *handlers.StudentHandler) {
    grp := r.Group("/api/students")
    {
        grp.POST("", h.Create)                     // POST   /api/students
        grp.GET("", h.List)                        // GET    /api/students
        grp.GET("/:student_id", h.Get)             // GET    /api/students/:student_id
        grp.PUT("/:student_id", h.Update)          // PUT    /api/students/:student_id
        grp.DELETE("/:student_id", h.Delete)       // DELETE /api/students/:student_id
        grp.GET("/search", h.Search)             // GET    /api/students/search
    }
}
