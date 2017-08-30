package main
import (
    "fmt"
    ht  "net/http"
)

func main(){
    ht.HandleFunc("/", ServerFunc)
    if err := ht.ListenAndServe("localhost:30001",nil); err != nil{
        fmt.Printf("Error :%v", err)
    }
}

func ServerFunc(w ht.ResponseWriter, req *ht.Request){
    fmt.Printf("%v \n %req", w, req)   
}
