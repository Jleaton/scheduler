package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Appointment struct {
	startTime time.Time
	endTime   time.Time
}

//NOTE: format of a time_slot is: 2006-01-02 15:04 +0000 UTC,60 where the last char is the duration in minutes
/** Business Logic:
1. Parse the timeSlot to dervie the start time and the duration from the given string
2. Take the start time string and convert it to a UTC time object
3. Using the Time.Add() method, we increment it by the duration in minutes
4. The result of step 3 is a end time object
5. Convert this back to a string that is in UTC format
6. Return back a appointment struct which contains the start and end time as strings in UTC format
*/
func (a *Appointment) DeriveStartAndEndTimeFromTimeSlot(timeSlot string) {

	subStrings := strings.Split(timeSlot, ",")

	startTimeString := subStrings[0]
	duration, err := strconv.Atoi(subStrings[1])
	if err != nil {
		fmt.Printf("failed deriving end time: %v", err)
	}

	startTime, err := time.Parse("2006-01-02 15:04 +0000 UTC", startTimeString)
	if err != nil {
		fmt.Printf("failed parsing time: %v", err)
	}

	endTime := startTime.Add(time.Minute * time.Duration(duration))

	a.startTime = startTime
	a.endTime = endTime
}
