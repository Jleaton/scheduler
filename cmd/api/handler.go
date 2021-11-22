package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func appointmentHandler(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	if len(params) == 0 {
		http.Error(w, "No path param provided", http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		timeSlot := params["time_slot"]

		if len(strings.Split(timeSlot, ",")) != 2 {
			http.Error(w, "Incorrect time_slot format", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, sch.IsTimeSlotAvailable(timeSlot))
	case http.MethodPost:

		timeSlot := params["time_slot"]

		if len(strings.Split(timeSlot, ",")) != 2 {
			http.Error(w, "Incorrect time_slot format", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, sch.BookTimeSlot(timeSlot))

	case http.MethodDelete:

		appointmentID := params["id"]

		if _, err := strconv.Atoi(appointmentID); err != nil {
			http.Error(w, "The appointment ID must be a integer", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, sch.CancelTimeSlot(appointmentID))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
