package main

import "fmt"
import "time"

var ch chan int

func main(){
    ch = make(chan int)
    go longRun(ch)
    time.Sleep(2e9)
    fmt.Println()
}

func longRun(ch chan int) {
    ch <- 10
}
