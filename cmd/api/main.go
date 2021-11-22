package main

import (
	"fmt"
	"net/http"

	"github.com/Jleaton/scheduler/internal/db"
	"github.com/Jleaton/scheduler/internal/service"
	"github.com/Jleaton/scheduler/pkg"
	"github.com/spf13/viper"
)

//provides global access of these properties to all files in the main package
var (
	sch *service.Scheduler
)

func envVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading in config file %v", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Printf("Error, the key needs to be a string %v", err)
	}

	return value
}

func main() {

	username := envVariable("USERNAME")
	password := envVariable("PASSWORD")
	host := envVariable("HOST")
	dbName := envVariable("DBNAME")
	ssl := envVariable("SSL")

	conn, err := pkg.DBConnect(username, password, host, dbName, ssl)
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
