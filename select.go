package main

import "fmt"
import "time"

func pump(ch chan int,n int){
    for i := 0;i < n; i++{
        ch <- i
    }
}

func main(){
    ch := make(chan int,5);
    ch2 := make(chan int, 10);
    pump(ch, 5)
    pump(ch2, 10)
    go func(){
    
        for{
            select{
                case b := <-ch :
                    fmt.Println(b)
                case c := <-ch2 :
                    fmt.Println(c)
            }
        }
    }()
    time.Sleep(1e9)
}
