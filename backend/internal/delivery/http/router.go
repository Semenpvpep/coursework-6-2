package http

import (
	"backend/internal/repository"
	"backend/internal/usecase"
	"backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo)
	appointmentUseCase := usecase.NewAppointmentUseCase(appointmentRepo)

	// Initialize handlers
	authHandler := NewAuthHandler(authUseCase)
	appointmentHandler := NewAppointmentHandler(appointmentUseCase)

	// Auth routes
	r.POST("/api/register", authHandler.Register)
	r.POST("/api/login", authHandler.Login)

	// Protected routes
	authMiddleware := middleware.AuthMiddleware()
	api := r.Group("/api")
	api.Use(authMiddleware)
	{
		api.POST("/appointments", appointmentHandler.CreateAppointment)
		api.GET("/appointments", appointmentHandler.GetAllAppointments)
		api.GET("/appointments/:id", appointmentHandler.GetAppointmentByID)
		api.PUT("/appointments/:id", appointmentHandler.UpdateAppointment)
		api.DELETE("/appointments/:id", appointmentHandler.DeleteAppointment)
	}

	// Serve frontend static files
	r.Static("/static", "./frontend")
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	r.GET("/dashboard", func(c *gin.Context) {
		c.File("./frontend/dashboard.html")
	})

	return r
}
