package entity

import (
	"time"
)

type Appointment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	DoctorName  string    `gorm:"not null" json:"doctorName"`
	PatientName string    `gorm:"not null" json:"patientName"`
	StartTime   time.Time `gorm:"not null" json:"startTime"`
	EndTime     time.Time `gorm:"not null" json:"endTime"`
	UserID      uint      `gorm:"not null" json:"-"`
}

type AppointmentRequest struct {
	Title       string    `json:"title" binding:"required"`
	DoctorName  string    `json:"doctorName" binding:"required"`
	PatientName string    `json:"patientName" binding:"required"`
	StartTime   time.Time `json:"startTime" binding:"required"`
	EndTime     time.Time `json:"endTime" binding:"required"`
}
