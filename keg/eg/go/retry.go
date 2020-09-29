package main

import (
	"fmt"
	"time"
)

func tryDo() error {
	// [Note] do whatever I want
	fmt.Println("tryDo encounter some problem...")
	return fmt.Errorf("do failed")
}

func retry1() {
	trySuccess := false
	for retries := 0; retries < 3; retries++ {
		err := tryDo()
		if err != nil {
			// wait a while before next try
			time.Sleep(time.Second)
		} else {
			// try success
			trySuccess = true
			break
		}
	}

	if !trySuccess {
		fmt.Println("Try failed")
	}
}

func retry2() {
	trySuccess := false
	for retries := 0; retries < 3; retries++ {
		err := tryDo()
		if err == nil {
			// try success
			trySuccess = true
			break
		}
		// wait a while before next try
		time.Sleep(time.Second)
	}

	if !trySuccess {
		fmt.Println("Try failed")
	}
}

func main() {
	//retry1()
	retry2()
}
