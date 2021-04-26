package main

import (
	"fmt"
	"sort"
)

func sortStrings() {
	strlist := []string{"hello", "how are you", "I'm fine"}
	sort.Strings(strlist)
	fmt.Println(strlist)
}

func sortSlice() {
	myList := []int64{1, 3, 7, 9, 8}
	sort.Slice(myList, func(i, j int) bool { return myList[i] < myList[j] })
	fmt.Println(myList)
}

func main() {
	//sortStrings()
	//sortSlice()
}
