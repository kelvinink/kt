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

func try() {
	today := time.Now()

	fmt.Println(today.Local())
	fmt.Println(time.Second)
}

func main() {
	//measureDuration()
	try()
}
