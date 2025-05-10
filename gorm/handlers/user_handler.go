package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm/models"
    "gorm/services"
)

type UserHandler struct {
    svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

func (h *UserHandler) Create(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.svc.Create(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, input)
}

func (h *UserHandler) List(c *gin.Context) {
    users, err := h.svc.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Get(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    user, err := h.svc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }
    c.JSON(http.StatusOK, user)
}
// PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := h.svc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }
    input.ID = uint(id)
    if err := h.svc.Update(&input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, input)
 }

// DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) { 
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    user, err := h.svc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }
    if err := h.svc.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
 }
 
 
// POST /login
func (h *UserHandler) Login(c *gin.Context) {
    var cred struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&cred); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
        return
    }

    // 1) Busca el usuario por email
    u, err := h.svc.GetByEmail(cred.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
        return
    }
    // 2) Comprueba que exista y que la contraseña coincida
    if u == nil || u.Password != cred.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales incorrectas"})
        return
    }

    // 3) Autenticación exitosa
    c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
