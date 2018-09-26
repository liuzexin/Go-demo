package main

import "fmt"

func main() {

	val := map[string]interface{}{
		"test": nil,
	}
	_, ok := val["test"]
	fmt.Println(ok)
}
