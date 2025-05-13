// routes/student_routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "gorm/handlers"
)

func SetupStudentRoutes(rg *gin.RouterGroup, h *handlers.StudentHandler) {
    // Aquí NO volvemos a poner "api", solo "students"
    grp := rg.Group("/students")
    {
        grp.POST("",           h.Create)    // POST   /api/students
        grp.GET("",            h.List)      // GET    /api/students
        grp.GET("/:id",        h.Get)       // GET    /api/students/:id
        grp.PUT("/:id",        h.Update)    // PUT    /api/students/:id
        grp.DELETE("/:id",     h.Delete)    // DELETE /api/students/:id
        grp.GET("/search",     h.Search)    // GET    /api/students/search
    }
}
