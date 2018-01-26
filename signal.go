package main

import (
    "os/signal"
    "time"
)

func main (){
    signal.Ignore()
    time.Sleep(30e9)
}
