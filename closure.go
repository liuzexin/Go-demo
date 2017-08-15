package main

import "fmt"

func main (){
    c := test("test")
    res := c(3, 4)
    fmt.Println(res)
}

func test(s string) func(int,int) int {
    return func (x int, y int) (z int){
        fmt.Println(s)       
        z = x + y
        return 
    }
}
