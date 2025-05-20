package http

import (
	"backend/internal/entity"
	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AppointmentHandler struct {
	appointmentUseCase usecase.AppointmentUseCase
}

func NewAppointmentHandler(appointmentUseCase usecase.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{appointmentUseCase: appointmentUseCase}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var request entity.AppointmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment := &entity.Appointment{
		Title:       request.Title,
		DoctorName:  request.DoctorName,
		PatientName: request.PatientName,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		UserID:      userID,
	}

	if err := h.appointmentUseCase.CreateAppointment(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	appointments, err := h.appointmentUseCase.GetAllAppointments(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (h *AppointmentHandler) GetAppointmentByID(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	appointment, err := h.appointmentUseCase.GetAppointmentByID(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var request entity.AppointmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment := &entity.Appointment{
		ID:          uint(id),
		Title:       request.Title,
		DoctorName:  request.DoctorName,
		PatientName: request.PatientName,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		UserID:      userID,
	}

	if err := h.appointmentUseCase.UpdateAppointment(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	if err := h.appointmentUseCase.DeleteAppointment(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
