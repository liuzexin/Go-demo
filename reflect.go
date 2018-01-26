package main

import (
    "fmt"
    "reflect"
)

type Person struct{
    name string
    age int
}

func main(){
    person := &Person{"Lzx", 10}
    value := reflect.ValueOf(*person)
    for i := 0; i < value.NumField(); i++{
        fmt.Printf("Field %d: %v\n", i, value.Field(i))
        fmt.Printf("Field %d: %v\n", i, value.Field(i).Type())
        fmt.Printf("Field %d: %v\n", i, value.Field(i).Interface())
    }
}
