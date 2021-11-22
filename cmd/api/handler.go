package main

import (
	"fmt"
	"net/http"
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

		if validatePathParameter(timeSlot, w) {
			fmt.Fprint(w, sch.IsTimeSlotAvailable(timeSlot))
		} else {
			return
		}

	case http.MethodPost:

		timeSlot := params["time_slot"]

		if validatePathParameter(timeSlot, w) {
			fmt.Fprint(w, sch.BookTimeSlot(timeSlot))
		} else {
			return
		}

	case http.MethodDelete:

		timeSlot := params["time_slot"]

		if validatePathParameter(timeSlot, w) {
			fmt.Fprint(w, sch.CancelTimeSlot(timeSlot))
		} else {
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func validatePathParameter(timeSlot string, w http.ResponseWriter) bool {

	isValid := false

	if len(strings.Split(timeSlot, ",")) == 2 {
		isValid = true
	} else {
		http.Error(w, "Incorrect time_slot format", http.StatusBadRequest)
	}

	return isValid
}
