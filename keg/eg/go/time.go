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

type Person struct {
	name string
	age int
}

func (p Person) sayHello ()  {
	fmt.Println("hi")
}

func try() {
	var man *Person
	man.sayHello()
}

func main() {
	//measureDuration()
	try()
}
