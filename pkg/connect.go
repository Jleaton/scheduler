package pkg

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

//handles the DB connection logic with the use of some helper functions connectionStringFormat() and driverConnect()
func DBConnect(username, password, host, dbName, sslMode string) (*sql.DB, error) {

	connectionString, err := connectionStringFormat(username, password, host, dbName, sslMode)
	if err != nil {
		return nil, err
	}

	db, err := driverConnect(connectionString)
	if err != nil {
		return nil, err
	}

	//ensure database connection is still active
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectionStringFormat(username, password, host, dbName, sslMode string) (string, error) {
	switch true {
	case username == "" || password == "" || host == "" || dbName == "" || sslMode == "":
		return "", fmt.Errorf("all parameters are required to have at least one character to DBConnect")
	case strings.Contains(username, " ") || strings.Contains(password, " ") || strings.Contains(host, " ") || strings.Contains(dbName, " ") || strings.Contains(sslMode, " "):
		return "", fmt.Errorf("can not have any space in parameters passed to DBConnect")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=%s", username, password, host, dbName, sslMode), nil
}

//connects to postgres driver with given connection string
func driverConnect(connectionStr string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}
	return db, err
}
