package main

import "fmt"
import "time"

func main(){
    go  longRun()
    time.Sleep(5)
    fmt.Println()
}

func longRun() int{
    time.Sleep(2)
    fmt.Println("sleep 10")
    return 10;
}
