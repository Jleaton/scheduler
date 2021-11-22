package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/appointment/{time_slot}", appointmentHandler).Methods(http.MethodGet, http.MethodPost, http.MethodDelete)

	return router
}
