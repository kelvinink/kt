package main

import (
	"fmt"
	"time"
)

func measureDuration() {
	var (
		start = time.Now()
	)

	// [Note] Do what ever you want here

	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
}

func parseInLocation() {
	const (
		DATE_FORMAT = "2006-01-02"
	)
	var (
		SHANGHAI_TZ, _ = time.LoadLocation("Asia/Shanghai")
	)

	t, err := time.ParseInLocation(DATE_FORMAT, "2020-12-27", SHANGHAI_TZ)
	if err != nil {
		fmt.Println("parse in location error")
		return
	}

	fmt.Println(t)
}

func yesterday() {
	var (
		t                = time.Now()
		year, month, day = t.Date()
	)

	// current time of yesterday
	yesterday := t.AddDate(0, 0, -1)
	fmt.Println(yesterday)

	// start of yesterday
	start_yesterday := time.Date(year, month, day-1, 0, 0, 0, 0, t.Location())
	fmt.Println(start_yesterday)

	// end of yesterday
	end_yesterday := start_yesterday.AddDate(0, 0, 1).Add(time.Nanosecond * -1)
	fmt.Println(end_yesterday)
}

func month() {
	var (
		t = time.Now()
	)

	// first day of month
	firstday := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	// last day of month
	lastday := firstday.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	fmt.Println(firstday)
	fmt.Println(lastday)
}

func main() {
	// measureDuration()
	// parseInLocation()
	// yesterday()
	// month()
}
