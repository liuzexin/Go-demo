package main

import "fmt"

type MyTest struct{
    c float64
}

type MyStruct struct{
    a,b int
    MyTest
}

func (this * MyStruct)echo() {
    fmt.Printf("%d %d", this.a, this.b)
}

func (this * MyTest)echo(){
    fmt.Printf("%f", this.c)
}

func main(){
    obj := &MyStruct{1, 2, MyTest{3}}
    obj.echo()   
    obj.MyTest.echo()
}
