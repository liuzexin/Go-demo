package main

import "fmt"

type MyStruct struct {
	a int
}

func main() {

	val := new(MyStruct)

	if val == nil {
		fmt.Println("true")
	}
	test := map[int]int{}

	if test == nil {
		fmt.Println("true")
	}
}
