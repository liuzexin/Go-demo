package main
import "fmt"

type MyInterface interface{
    MyMethod() float64
}

type MyStruct struct{
    MyInterface
}

func (this *MyStruct) MyMethod() float64{
    fmt.Println("test");
    return 64.3
}


func main(){
    
    var myVar MyInterface 
    myVar = new(MyStruct)
    myVar.MyMethod()
}
