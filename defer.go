package main    
import "fmt"

func f() (ret int){
    defer func(){
        ret ++
    }()

    return 1
}

func main(){
    res := f()
    fmt.Println(res)
}
