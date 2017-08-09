package main

import "os"
import "fmt"

func main(){
    if _,err := os.Open("/notexsit/path");err != nil {
        fmt.Println(err);
    }
    switchFunc()
}

func switchFunc(){
    var a = 1
    var b = 2
    switch{
        case a > b :
            fmt.Println("max is", a)
        case b > a :
            fmt.Println("max is", b)
    }
}
