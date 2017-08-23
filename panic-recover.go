package main

import "fmt"

func main(){

    defer func(){
        fmt.Println("Dangerous")
        if err := recover(); err != nil{
            fmt.Println("Do Something")
        }
        panic("Sorry I must exit now")
        
    }()
    testPanic()
}

func testPanic(){
    panic("Error: FATAL ERROR-You can't call this funciton.")    
}
