package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	a int
	b int
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

	var c int
	if &c == nil {
		fmt.Println("true")
	}
	var s []int

	if s == nil {
		fmt.Println("--true")
	}

	val2 := MyStruct{}
	val3 := &val2
	if &val2 == val {
		fmt.Println("---------true")
	}
	if &val2 == val3 {
		fmt.Println("pointer--------true")
	}

	if reflect.DeepEqual(val, val2) {
		fmt.Println("true")
	}
	if val2 == *val {
		fmt.Println("----true")
	}

}
