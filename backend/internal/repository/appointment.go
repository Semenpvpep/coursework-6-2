package repository

import (
	"backend/internal/entity"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Create(appointment *entity.Appointment) error
	GetAll(userID uint) ([]entity.Appointment, error)
	GetByID(id uint, userID uint) (*entity.Appointment, error)
	Update(appointment *entity.Appointment) error
	Delete(id uint, userID uint) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Create(appointment *entity.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetAll(userID uint) ([]entity.Appointment, error) {
	var appointments []entity.Appointment
	if err := r.db.Where("user_id = ?", userID).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *appointmentRepository) GetByID(id uint, userID uint) (*entity.Appointment, error) {
	var appointment entity.Appointment
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *appointmentRepository) Update(appointment *entity.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *appointmentRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Appointment{}).Error
}
