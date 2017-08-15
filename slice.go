package main

import "fmt"

func main(){
    slice := []int{1,2,3,4,5,6,7,8,9}
    slice1 := slice[2:5]
    for _,val := range slice1{
        fmt.Println(val)
    }
    fmt.Println("--------------------------")

    slice1 = slice1[4:7]

    for _,val := range slice1{
        fmt.Println(val)
    }
}
