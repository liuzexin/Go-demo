package main
import "fmt"

type MyStruct struct{
    x int
    y int
}

func main() {
    var a []int
    var b = MyStruct{1,2}
    if b.x == 1{
        fmt.Println(b.x)
    }
    fmt.Printf("hello,word \n", a, b.x)
}
func anthoer(){
    var b = make([]int,3, 5)
    fmt.Printf("test arr\n", cap(b))
}
