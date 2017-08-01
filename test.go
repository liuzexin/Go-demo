package main
import "fmt"
import "net/http"
type MyStruct struct{
    x int
    y int
}
//test
/**
var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}
*/
type Api interface{
    api() int
}
func main() {
    var a []int
    var b = MyStruct{1,2}
    if b.x == 1{
        fmt.Println(b.x)
    }
    err := http.ListenAndServe("localhost:60000", b);
    if err != nil {
        fmt.Println(err)
    }
    other()
    second()
    third()
    sixth()
    c := &b
    c.fifth()
    fmt.Printf("hello,word \n", a, b.x)
    seventh()
    eighth()
}
func another(){
    var b = make([]int,3, 5)
    fmt.Printf("test arr\n", cap(b))
}

var m map[string]MyStruct

func other(){
    m = make(map[string]MyStruct)
    m["test"] = MyStruct{
        10,
        20,
    }
    fmt.Println(m["test"])
}
func second(){
    testMap := make([]int, 10)//创建map需要使用此函数
    for i := range testMap{
        fmt.Printf("%d,%d \n",i,  testMap[i])
    }
}

func third(){
    test := func (x int) float64{
        return float64(x) * 3.0
    }
    res := fourth(test)
    fmt.Println(res)
}

func fourth(f func (int) float64) float64{
    return f(3)
}

func (s *MyStruct) fifth() int{
    res := s.x + s.y
    fmt.Println(res)
    return res
}
func sixth(){
    var a int
    a = 10
    b := &a
    fmt.Println(a, b)
}
func seventh(){
    var a Api
    b := MyStruct{1, 2}
    a = &b
    c := a.api()
    fmt.Println(c)
}
func (m MyStruct) api() int{
    return 64
}
func eighth(){
    var a [4]byte
    a = [4]byte{1, 2, 3, 4}
    b := string(a[:])
    fmt.Println(b)
}
func (ms MyStruct ) ServeHTTP (w http.ResponseWriter ,r *http.Request){
    fmt.Fprint(w, "My Struct")
}
