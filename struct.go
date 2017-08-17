package main
import "fmt"

type T struct{
    a int
    b float64
}

func main(){
    var t *T = new(T)
    t.a = 100
    t.b = 1000
    fmt.Println(t)
}
