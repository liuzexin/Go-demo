package main

import "fmt"

func main(){
    v := 1
    if val, ok := v.(int){
        fmt.Println(ok)
    }
}
