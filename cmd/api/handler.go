package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jleaton/scheduler/internal/service"
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

		if ok := validatePathParameter(timeSlot, w); !ok {
			userResponse(w, "invalid time slot format", nil)
			return
		}

		isAvailable, err := sch.IsTimeSlotAvailable(timeSlot)
		if err != nil {
			if errors.Is(err, service.ErrNoRecordModified) {
				userResponse(w, "There was an error checking for avaialbility", nil)
			} else if errors.Is(err, service.ErrConflict) {
				userResponse(w, false, nil)
			} else {
				userResponse(w, http.StatusText(http.StatusInternalServerError), err)
			}
			return
		}

		userResponse(w, isAvailable, nil)
		return

	case http.MethodPost:

		timeSlot := params["time_slot"]

		if ok := validatePathParameter(timeSlot, w); !ok {
			userResponse(w, "invalid time slot format", nil)
			return
		}

		err := sch.BookTimeSlot(timeSlot)
		if err != nil {
			if errors.Is(err, service.ErrNoRecordModified) {
				userResponse(w, "appointment could not be booked", nil)
			} else if errors.Is(err, service.ErrConflict) {
				userResponse(w, "appointment could not be booked due to a conflict", nil)
			} else {
				userResponse(w, http.StatusText(http.StatusInternalServerError), err)
			}
			return
		}

		userResponse(w, "appointment booked", nil)

	case http.MethodDelete:
		timeSlot := params["time_slot"]

		if ok := validatePathParameter(timeSlot, w); !ok {
			userResponse(w, "invalid time slot format", nil)
			return
		}

		err := sch.CancelTimeSlot(timeSlot)
		if err != nil {
			if errors.Is(err, service.ErrNoRecordModified) {
				userResponse(w, "no record with the time given exists", nil)
			} else {
				userResponse(w, http.StatusText(http.StatusInternalServerError), err)
			}
			return
		}

		userResponse(w, "appointment cancelled", nil)
		return

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func userResponse(w http.ResponseWriter, data interface{}, err error) {

	if err != nil {
		errLogger.Println(err)
	}
	fmt.Fprint(w, data)
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
