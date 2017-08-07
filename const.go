package main

import "fmt"

const (
    Unknow = 0
    Female = 1
    male = 2
)

const (
    a = iota
    b
    c
)

func main(){
    var sex = 0
    switch (sex){
        case Unknow:
            fmt.Println("Unknow")
        case Female:
            fmt.Println("female")
        case male:
            fmt.Println("male")
    }

    fmt.Println(a, b, c)
}
