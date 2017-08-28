package main

import "fmt"
import "time"
type Empty interface {}

type semaphore chan Empty

var a int = 1

func main(){
    n := 1
    var sem semaphore = make(semaphore, n)
    for i := 0; i < 10; i++{
        go countPlus(sem)
    }
    time.Sleep(300000)
    fmt.Println(a)    
}

func countPlus(sem  semaphore){
    sem.lock()
    a ++
    sem.unlock()
}




func (sem semaphore)lock(){
    e := new(Empty)
    for i := 0;i < 1; i++{
        sem <- e
    }
}

func (sem semaphore)unlock(){
    for i:=0;i < 1;i++{
        <- sem
    }
}
