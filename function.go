package main
import "fmt"

func main(){
    res := test(1, 2)
    fmt.Println(res)
}

func test(y int, x int)( z int){
    z = x + y
    return // directly return with no variable
}
