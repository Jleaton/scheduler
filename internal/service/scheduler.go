package service

import (
	"errors"
	"fmt"

	"github.com/Jleaton/scheduler/internal/db"
)

var (
	ErrNoRecordModified = errors.New("no record modified or inserted")
	ErrConflict         = errors.New("appointment time is already booked")
)

type Scheduler struct {
	dbRepo *db.DB
}

func New(dbRepo *db.DB) *Scheduler {
	return &Scheduler{
		dbRepo: dbRepo,
	}
}

func (s *Scheduler) IsTimeSlotAvailable(timeSlot string) (bool, error) {

	var appointment db.Appointment
	appointment.DeriveStartAndEndTimeFromTimeSlot(timeSlot)

	isAvailable, err := s.dbRepo.SelectAppointmentAvailablility(&appointment)
	if err != nil {
		return false, fmt.Errorf("failed to check appointment availability: %w", err)
	}

	return isAvailable, nil
}

func (s *Scheduler) BookTimeSlot(timeSlot string) error {

	var appointment db.Appointment
	appointment.DeriveStartAndEndTimeFromTimeSlot(timeSlot)

	rowsAffected, err := s.dbRepo.InsertAppointment(&appointment)
	if err != nil {
		return fmt.Errorf("database failed to book appointment: %w", err)
	}

	if rowsAffected != 1 {
		return ErrNoRecordModified
	}

	if rowsAffected != 1 {
		return ErrConflict
	}

	return nil
}

func (s *Scheduler) CancelTimeSlot(timeSlot string) error {

	var appointment db.Appointment
	appointment.DeriveStartAndEndTimeFromTimeSlot(timeSlot)

	rowsAffected, err := s.dbRepo.CancelAppointment(&appointment)
	if err != nil {
		return fmt.Errorf("database failed to cancel appointment: %w", err)
	}

	if rowsAffected != 1 {
		return ErrNoRecordModified
	}

	return nil
}
