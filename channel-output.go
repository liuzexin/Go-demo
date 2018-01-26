package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	in(ch)
	out(ch)
}

func in(ch chan<- int) {
	ch <- 1
}

func out(ch <-chan int) {
	fmt.Println(<-ch)
}
