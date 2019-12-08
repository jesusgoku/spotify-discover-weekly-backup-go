package main

import (
	"time"

	"github.com/snabb/isoweek"
)

// GetFirstDayOfWeek return first day of week
func GetFirstDayOfWeek(inputTime time.Time) time.Time {
	year, week := inputTime.ISOWeek()

	return isoweek.StartTime(year, week, inputTime.Location())
}

// GetFirstDayForCurrentWeek return first day of current week
func GetFirstDayForCurrentWeek() time.Time {
	return GetFirstDayOfWeek(time.Now())
}
