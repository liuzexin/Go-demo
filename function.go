package main
import "fmt"

func main(){
    res := test(1, 2)
    fmt.Println(res)
    var arr []int
    arr=[]int{1, 3, 4}
    argTest(arr...)
    argTest(1, 2, 3)
}

func test(y int, x int)( z int){
    z = x + y
    return // directly return with no variable
}

func argTest(arr ...int){
    for _,val := range arr{
        fmt.Println(val)    
    }
}
