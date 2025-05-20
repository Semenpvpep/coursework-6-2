package http

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authUseCase.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var credentials entity.AuthRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Добавьте эти строки для очистки пробелов
	credentials.Login = strings.TrimSpace(credentials.Login)
	credentials.Password = strings.TrimSpace(credentials.Password)

	response, err := h.authUseCase.Login(&credentials)
	log.Printf("Login failed for user %s: %v", credentials.Login, err) // Логируем ошибку
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "details": err.Error()}) // Возвращаем детали клиенту
		return
	}

	c.JSON(http.StatusOK, response)
}
