package main

import (
	"fmt"
)

func createSet() {
	mySet := make(map[string]bool)
	mySet["alice"] = true
	fmt.Println(mySet)
}

func insertSet() {
	mySet := make(map[string]bool)
	mySet["alice"] = true
	fmt.Println(mySet)
	mySet["bob"] = true
	fmt.Println(mySet)
}

func deleteSet() {
	mySet := make(map[string]bool)
	mySet["alice"] = true
	fmt.Println(mySet)
	mySet["bob"] = true
	fmt.Println(mySet)
	delete(mySet, "alice")
	fmt.Println(mySet)
}

func setMember() {
	mySet := make(map[string]bool)
	mySet["alice"] = true
	_, in := mySet["alice"]
	fmt.Println(in)
}

func setSize() {
	mySet := make(map[string]bool)
	mySet["alice"] = true
	mySet["bob"] = true
	mySet["carol"] = true
	size := len(mySet)
	fmt.Println(size)
}

func setIterate() {
	mySet := make(map[string]bool)
	mySet["alice"] = true
	mySet["bob"] = true
	mySet["carol"] = true
	for k := range mySet {
		fmt.Println(k)
	}
}

func main() {
	//createSet()
	//insertSet()
	//deleteSet()
	//setMember()
	//setSize()
	//setIterate()
}
