package main

import (
	"encoding/json"
	"fmt"
)

func marshal() {
	myMap := map[string]int{
		"age":    18,
		"weight": 55,
		"height": 170,
	}

	// marshal
	bytes, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("Json marshal error")
	}

	fmt.Println(bytes)

	// unmarshal
	err = json.Unmarshal(bytes, &myMap)
	if err != nil {
		fmt.Println("Json unmarshal error")
	}

	fmt.Println(myMap)
}

func main() {
	marshal()
}
