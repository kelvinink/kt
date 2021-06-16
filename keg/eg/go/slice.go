package main

import "fmt"

func modifySlice(s []int) {
	for i := range s {
		s[i] = s[i] * 2
	}
}

func modifyArray(arr [5]int) {
	for i := range arr {
		arr[i] = arr[i] * 2
	}
}

// Compare the difference of array and slice modification
func compareArraySliceModification() {
	mySlice := []int{1, 2, 3, 4, 5}
	myArray := [5]int{1, 2, 3, 4, 5}
	modifySlice(mySlice)
	modifyArray(myArray)
	fmt.Printf("Slice after mofify:%v", mySlice)
	fmt.Println()
	fmt.Printf("Array after mofify:%v", myArray)
}

func main() {
	compareArraySliceModification()
}
