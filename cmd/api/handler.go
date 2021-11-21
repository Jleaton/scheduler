package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jleaton/scheduler/internal/db"
	"github.com/gorilla/mux"
)

func appointmentAvailabilityHandler(w http.ResponseWriter, req *http.Request) {

	var appointment db.Appointment

	err := json.NewDecoder(req.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprint(w, sch.IsTimeSlotAvailable(&appointment))
}

func bookAppointmentHandler(w http.ResponseWriter, req *http.Request) {

	var appointment db.Appointment

	err := json.NewDecoder(req.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, sch.BookTimeSlot(&appointment))
}

func appointmentCancelHandler(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	if len(params) == 0 {
		http.Error(w, "No ID param provided", http.StatusBadRequest)
		return
	}

	appointmentID := params["id"]

	if _, err := strconv.Atoi(appointmentID); err != nil {
		http.Error(w, "The appointment ID must be a integer", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, sch.CancelTimeSlot(appointmentID))
}
