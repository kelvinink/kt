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

	bytes, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("Json marshal error")
	}

	fmt.Println(bytes)
}

func unmarshal() {
	myMap := map[string]int{
		"age":    18,
		"weight": 55,
		"height": 170,
	}

	str := `{"age":18,"height":170,"weight":55}`
	err := json.Unmarshal([]byte(str), &myMap)
	if err != nil {
		fmt.Println("Json unmarshal error")
	}

	fmt.Println(myMap)
}

func marshalStruct() {
	type Person struct {
		Name *string `json:"name"`
		Age  *int64  `json:"age"`
	}

	age := int64(25)
	name := "kelvin"

	person := Person{
		Name: &name,
		Age:  &age,
	}

	bytes, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("Json marshal error")
	}

	print(string(bytes))
}

func main() {
	//marshal()
	//unmarshal()
	//marshalStruct()
}
