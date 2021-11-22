package service

import (
	"fmt"

	"github.com/Jleaton/scheduler/internal/db"
)

type Scheduler struct {
	dbRepo *db.DB
}

func New(dbRepo *db.DB) *Scheduler {
	return &Scheduler{
		dbRepo: dbRepo,
	}
}

func (s *Scheduler) IsTimeSlotAvailable(timeSlot string) bool {

	var appointment db.Appointment
	appointment.DeriveStartAndEndTimeFromTimeSlot(timeSlot)

	isAvailable, err := s.dbRepo.SelectAppointmentAvailablility(&appointment)
	if err != nil {
		fmt.Printf("failed to check appointment availability: %v", err)
	}

	return isAvailable
}

func (s *Scheduler) BookTimeSlot(timeSlot string) string {

	var appointment db.Appointment
	appointment.DeriveStartAndEndTimeFromTimeSlot(timeSlot)

	message, err := s.dbRepo.InsertAppointment(&appointment)
	if err != nil {
		fmt.Printf("failed to insert appointment: %v", err)
	}

	return message
}

func (s *Scheduler) CancelTimeSlot(appointmentID string) string {

	message, err := s.dbRepo.CancelAppointment(appointmentID)
	if err != nil {
		fmt.Printf("failed to cancel appointment: %v", err)
	}

	return message
}
