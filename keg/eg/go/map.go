package main

import "fmt"

func createMap() {
	// alt 1
	myMap1 := make(map[string]int)

	// alt 2
	// [Note] 注意myMap2为nil
	var myMap2 map[string]int

	// alt 3
	myMap3 := map[string]int{
		"age":    18,
		"weight": 55,
		"height": 170,
	}

	myMap1["age"] = 26
	myMap2["age"] = 26 // [Note] panic
	myMap3["age"] = 26

	fmt.Println(myMap1)
	fmt.Println(myMap2)
	fmt.Println(myMap3)
}

func insertMap() {
	myMap := make(map[string]int)
	myMap["age"] = 26

	fmt.Println(myMap)
}

func deleteMap() {
	myMap := make(map[string]int)
	myMap["age"] = 26

	fmt.Println(myMap)
	delete(myMap, "age")
	fmt.Println(myMap)
}

func mapMember() {
	myMap := map[string]int{
		"age":    18,
		"weight": 55,
		"height": 170,
	}

	member, ok := myMap["age"]
	fmt.Println(member, ok)
}

func mapSize() {
	myMap := map[string]int{
		"age":    18,
		"weight": 55,
		"height": 170,
	}

	size := len(myMap)
	fmt.Println(size)
}

func mapIterate() {
	myMap := map[string]int{
		"age":    18,
		"weight": 55,
		"height": 170,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

func main() {
	//createMap()
	//insertMap()
	//deleteMap()
	//mapMember()
	//mapSize()
	//mapIterate()
}
