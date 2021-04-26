package main

import (
	"fmt"
	"time"
)

func measureDuration() {
	start := time.Now()

	// [Note] Do what ever you want here

	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
}

func parseInLocation() {
	const DATE_FORMAT = "2006-01-02"
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

func main() {
	//measureDuration()
	//parseInLocation()
}
