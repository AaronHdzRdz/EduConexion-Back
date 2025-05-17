package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm/models"
    "gorm/services"
)

type SubjectHandler struct {
    svc *services.SubjectService
}

func NewSubjectHandler(svc *services.SubjectService) *SubjectHandler {
    return &SubjectHandler{svc: svc}
}

func (h *SubjectHandler) Create(c *gin.Context) {
    var input models.Subject
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // extraer userID del contexto
    uid, _ := c.Get("userID")
    input.UserID = uid.(uint)

    if err := h.svc.Create(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, input)
}

func (h *SubjectHandler) List(c *gin.Context) {
    uid, _ := c.Get("userID")
    list, err := h.svc.GetAllByUser(uid.(uint))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, list)
}

func (h *SubjectHandler) Get(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    uid, _ := c.Get("userID")
    sub, err := h.svc.GetByID(uint(id), uid.(uint))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if sub == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
        return
    }
    c.JSON(http.StatusOK, sub)
}

func (h *SubjectHandler) Update(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    var input models.Subject
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    uid, _ := c.Get("userID")
    existing, err := h.svc.GetByID(uint(id), uid.(uint))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if existing == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
        return
    }
    // actualiza campos
    existing.Name = input.Name
    existing.Code = input.Code
    if err := h.svc.Update(existing); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, existing)
}

func (h *SubjectHandler) Delete(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    uid, _ := c.Get("userID")
    if err := h.svc.Delete(uint(id), uid.(uint)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
