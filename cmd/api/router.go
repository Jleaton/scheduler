package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/appointment/availability/{time_slot}", appointmentHandler).Methods(http.MethodGet)
	router.HandleFunc("/appointment/book/{time_slot}", appointmentHandler).Methods(http.MethodPost)
	router.HandleFunc("/appointment/cancel/{id}", appointmentHandler).Methods(http.MethodDelete)

	return router
}
