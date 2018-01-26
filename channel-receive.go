package main

import (
//"fmt"
)

func main() {
	ch := make(chan int, 1)
	in(ch)
}

func in(ch <-chan int) {
	ch <- 1
}
