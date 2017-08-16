package main
import "fmt"

func main(){
    str := "测试字符串"
    for i,v := range str{
        fmt.Printf("%d------%c",i,v)
    }
}
