package main

import (
	"fmt"
	"net/http"

	"github.com/Jleaton/scheduler/internal/db"
	"github.com/Jleaton/scheduler/internal/service"
)

//provides global access of these properties to all files in the main package
var (
	sch *service.Scheduler
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		fmt.Printf("failed to create db connection: %v", err)
	}

	dbRepo := db.New(conn)

	sch = service.New(dbRepo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: Router(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to start server: %v", err)
	}

}
