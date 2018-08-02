package shutdown

import (
	"testing"
	"time"
)

func getTimeNow(hour int, min int) time.Time {
	timeNow := time.Now()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, min, 0, 0, time.Local)
}

func TestShouldShutdown(t *testing.T) {
	type suiltTests struct {
		timeRange    string
		timeNow      time.Time
		expectResult bool
	}

	seriesTests := []suiltTests{
		{timeRange: "23:00-07:00", timeNow: getTimeNow(23, 00), expectResult: true},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(23, 01), expectResult: true},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(06, 59), expectResult: true},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(07, 00), expectResult: true},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(22, 59), expectResult: false},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(07, 01), expectResult: false},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(12, 00), expectResult: false},
		{timeRange: "23:00-07:00", timeNow: getTimeNow(20, 00), expectResult: false},

		{timeRange: "08:00-07:00", timeNow: getTimeNow(8, 00), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(8, 01), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(6, 59), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(7, 00), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(23, 00), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(00, 00), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(01, 00), expectResult: true},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(7, 01), expectResult: false},
		{timeRange: "08:00-07:00", timeNow: getTimeNow(7, 59), expectResult: false},

		{timeRange: "07:00-08:00", timeNow: getTimeNow(7, 00), expectResult: true},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(7, 01), expectResult: true},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(7, 59), expectResult: true},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(8, 00), expectResult: true},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(6, 59), expectResult: false},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(8, 01), expectResult: false},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(23, 00), expectResult: false},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(00, 00), expectResult: false},
		{timeRange: "07:00-08:00", timeNow: getTimeNow(01, 00), expectResult: false},
	}

	for _, serie := range seriesTests {
		if serie.expectResult != Should(serie.timeNow, serie.timeRange) {
			t.Errorf("TimeRange [%s] and TimeNow [%s] result should be [%t]", serie.timeRange, serie.timeNow, serie.expectResult)
		}
	}
}
