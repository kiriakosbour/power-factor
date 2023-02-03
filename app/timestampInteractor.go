package app

import (
	"log"
	"power-factor/domain"
	"time"
)

const (
	layout = "20060102T150405Z"
)

type TimestampInteractor struct {
}

func NewTimestampInteractor() *TimestampInteractor {
	return &TimestampInteractor{}
}
func (t TimestampInteractor) TimestampService(tms domain.Timestamp) ([]string, domain.ErrorResp) {
	intervals := []string{"1h", "1d", "1mo", "1y"}

	zone, err := time.LoadLocation(tms.Timezone)
	if err != nil {
		log.Printf("Error from Location %s", err.Error())
		return nil, domain.ErrorResp{
			Status: "error",
			Desc:   "illegal location provided",
		}
	}
	startTimestamp, err := time.Parse(layout, tms.Start)
	if err != nil {
		log.Printf("Error from Parsing start date %s", err.Error())
		return nil, domain.ErrorResp{
			Status: "error",
			Desc:   "illegal start date provided",
		}
	}
	endTimestamp, err := time.Parse(layout, tms.End)
	if err != nil {
		log.Printf("Error from Parsing end date %s", err.Error())
		return nil, domain.ErrorResp{
			Status: "error",
			Desc:   "illegal end date provided",
		}
	}
	if endTimestamp.Before(startTimestamp) {
		log.Printf("Illegal start date %s", startTimestamp.GoString())
		return nil, domain.ErrorResp{
			Status: "error",
			Desc:   "start date cannot be before end date",
		}
	}
	if !contains(intervals, tms.Interval) {
		return nil, domain.ErrorResp{
			Status: "error",
			Desc:   "unsupported period",
		}
	}
	tstamp := startTimestamp.Add(time.Hour)
	tstamp = tstamp.In(zone)
	endTimestamp = endTimestamp.In(zone)
	tstamp = truncateToHour(tstamp)
	var intervalTimestamps []string
	for tstamp.Before(endTimestamp) {
		if tms.Interval == "1mo" {
			if lastDateOfMonth(tstamp).Before(endTimestamp) {
				intervalTimestamps = append(intervalTimestamps, lastDateOfMonth(tstamp).In(zone).UTC().Format(layout))
			}
		} else if tms.Interval == "1y" {
			if lastDateOfYear(tstamp).Before(endTimestamp) {
				intervalTimestamps = append(intervalTimestamps, lastDateOfYear(tstamp).In(zone).UTC().Format(layout))
			}
		} else {

			intervalTimestamps = append(intervalTimestamps, tstamp.In(zone).UTC().Format(layout))
		}
		tstamp = addInterval(tms.Interval, tstamp)

	}

	return intervalTimestamps, domain.ErrorResp{}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func addInterval(interval string, timestamp time.Time) time.Time {

	switch interval {
	case "1h":
		timestamp = timestamp.Add(time.Hour)
	case "1d":
		timestamp = timestamp.AddDate(0, 0, 1)
	case "1mo":
		timestamp = timestamp.AddDate(0, 1, 0)
	case "1y":
		timestamp = timestamp.AddDate(1, 0, 0)
	}
	return timestamp
}
func truncateToHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func lastDateOfMonth(t time.Time) time.Time {
	tsYear, tsMonth, _ := t.Date()
	firstOfMonth := time.Date(tsYear, tsMonth, 1, t.Hour(), 0, 0, 0, t.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return lastOfMonth
}

func lastDateOfYear(t time.Time) time.Time {
	tsYear, _, _ := t.Date()
	lastOfYear := time.Date(tsYear, 12, 31, t.Hour(), 0, 0, 0, t.Location())
	return lastOfYear
}
