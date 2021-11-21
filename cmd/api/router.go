package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/appointment/availability", appointmentAvailabilityHandler).Methods(http.MethodGet)
	router.HandleFunc("/appointment/book", bookAppointmentHandler).Methods(http.MethodPost)
	router.HandleFunc("/appointment/cancel/{id}", appointmentCancelHandler).Methods(http.MethodDelete)

	return router
}
