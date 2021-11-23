package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func (db *DB) SelectAppointmentAvailablility(appointment *Appointment) (bool, error) {

	var rowCount = 0

	sqlStatement := `SELECT COUNT(*) FROM appointments WHERE
( appointments.start_time > $1 AND appointments.start_time < $2 )
OR 
( appointments.end_time > $1 AND  appointments.end_time < $2 ) 
OR
(appointments.start_time <= $1 AND appointments.end_time >= $2)`

	res := db.conn.QueryRow(sqlStatement, appointment.startTime, appointment.endTime)

	err := res.Scan(&rowCount)
	if err != nil {
		return false, fmt.Errorf("failed to query db %w", err)
	}

	if rowCount == 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (db *DB) InsertAppointment(appointment *Appointment) (int, error) {

	//NOTE: filtering by day not completed
	sqlStatement := `INSERT INTO appointments (start_time, end_time) 
						SELECT $1, $2 WHERE NOT EXISTS (
							SELECT 1 FROM appointments WHERE 
								(appointments.start_time > $1 AND appointments.start_time < $2) OR 
								(appointments.end_time > $1 AND  appointments.end_time < $2) OR
								(appointments.start_time <= $1 AND appointments.end_time >= $2)
							)`

	res, err := db.conn.Exec(sqlStatement, appointment.startTime, appointment.endTime)
	if err != nil {
		return 0, fmt.Errorf("failed to execute transaction on the database: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to calculate rows affected in transaction on the database: %w", err)
	}

	return int(count), nil
}

func (db *DB) CancelAppointment(appointment *Appointment) (int, error) {

	sqlStatement := `DELETE FROM appointments WHERE start_time = $1 AND end_time = $2;`
	res, err := db.conn.Exec(sqlStatement, appointment.startTime, appointment.endTime)
	if err != nil {
		return 0, fmt.Errorf("failed to execute transaction on the database: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to calculate rows affected in transaction on the database: %w", err)
	}

	return int(count), nil
}
