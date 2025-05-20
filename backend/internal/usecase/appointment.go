package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
)

type AppointmentUseCase interface {
	CreateAppointment(appointment *entity.Appointment) error
	GetAllAppointments(userID uint) ([]entity.Appointment, error)
	GetAppointmentByID(id uint, userID uint) (*entity.Appointment, error)
	UpdateAppointment(appointment *entity.Appointment) error
	DeleteAppointment(id uint, userID uint) error
}

type appointmentUseCase struct {
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentUseCase(appointmentRepo repository.AppointmentRepository) AppointmentUseCase {
	return &appointmentUseCase{appointmentRepo: appointmentRepo}
}

func (uc *appointmentUseCase) CreateAppointment(appointment *entity.Appointment) error {
	return uc.appointmentRepo.Create(appointment)
}

func (uc *appointmentUseCase) GetAllAppointments(userID uint) ([]entity.Appointment, error) {
	return uc.appointmentRepo.GetAll(userID)
}

func (uc *appointmentUseCase) GetAppointmentByID(id uint, userID uint) (*entity.Appointment, error) {
	return uc.appointmentRepo.GetByID(id, userID)
}

func (uc *appointmentUseCase) UpdateAppointment(appointment *entity.Appointment) error {
	return uc.appointmentRepo.Update(appointment)
}

func (uc *appointmentUseCase) DeleteAppointment(id uint, userID uint) error {
	return uc.appointmentRepo.Delete(id, userID)
}
