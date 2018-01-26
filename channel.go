package main

import "fmt"
import "time"

func main() {
	ch := make(chan string)
	input(ch)
	print(ch)
	time.Sleep(1e9)
}

func input(ch chan string) {
	ch <- "10"
}

func print(ch chan string) {
	fmt.Println("run")
	fmt.Println(<-ch)
}
