package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(p)
	fmt.Println(os.Getenv("USER"))
}
