package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm/models"
    "gorm/services"
)

type StudentHandler struct {
    svc *services.StudentService
}

func NewStudentHandler(svc *services.StudentService) *StudentHandler {
    return &StudentHandler{svc: svc}
}

func (h *StudentHandler) Get(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("student_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    st, err := h.svc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if st == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
        return
    }
    c.JSON(http.StatusOK, st)
}

func (h *StudentHandler) Update(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("student_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    var input models.Student
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    existing, err := h.svc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if existing == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Alumno no encontrado"})
        return
    }
    input.ID = uint(id)
    if err := h.svc.Update(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, input)
}

func (h *StudentHandler) Delete(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("student_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    if err := h.svc.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *StudentHandler) Search(c *gin.Context) {
	query := c.Query("query")
	students, err := h.svc.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) Create(c *gin.Context) {
  var input models.Student
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  // extrae userID
  uid, _ := c.Get("userID")
  input.UserID = uid.(uint)

  if err := h.svc.Create(&input); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusCreated, input)
}

func (h *StudentHandler) List(c *gin.Context) {
  uid, _ := c.Get("userID")
  students, err := h.svc.GetAllByUser(uid.(uint))
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, students)
}