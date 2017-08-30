 package main

 import (
    "fmt"
    "net"
    "os"
 )

 func main(){
    fmt.Println("Start the tcp server")
    listener, err := net.Listen("tcp", "localhost:30000")
    checkError(err) 
    for {
        conn,err := listener.Accept()
        checkError(err)  
        go receive(conn)
    }
}

func checkError(err error){
    if err != nil{
        fmt.Println("Error:%v", err)
        os.Exit(250)
    }
}

func receive(conn net.Conn){
    for{
        buf := make([]byte, 8)
        len, err := conn.Read(buf)
        checkError(err)
        fmt.Printf("Recevied data:%v",string(buf[:len]))
    }
}
