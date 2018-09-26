package main

import "fmt"

var a float64
var b int8
var c int8

func init() {
	a = 10
	a := 3
	var b float64 = 3
	fmt.Println(a, b)
}

func main() {
	b = 10
	c = 10
	if b == c {
		fmt.Println("ok")
	}
}
