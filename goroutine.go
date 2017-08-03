package main
import (
    "fmt"
    //"time"
)
func main(){
    //go time.Sleep(10 * time.Second)
    ch := make(chan int ,10)
    go second(ch)
    for i := range ch{
        fmt.Println(i)
    }
    go second(ch)
    for {
        select{
            case  <-ch:
                fmt.Println(<-ch)
            default:
                fmt.Println("default")
        }
    }
}
func second(ch chan int){
    for i:= 10; i > 1; i--{
        ch <- i
    }
    close(ch)
}
