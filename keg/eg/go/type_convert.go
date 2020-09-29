package main

import (
	"fmt"
	"strconv"
)

func string2Other() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, f, i, u)
}

func other2String() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 10)
	s4 := strconv.FormatUint(42, 10)
	fmt.Println(s1, s2, s3, s4)
}

func main() {
	string2Other()
	other2String()
}
